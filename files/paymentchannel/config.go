package paymentchannel

import (
	"fmt"
	"os"
	"restapidemo/mdbal"
	"sync"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	ewallet "perun.network/go-perun/backend/ethereum/wallet"

	plog "perun.network/go-perun/log"
	plogrus "perun.network/go-perun/log/logrus"

	"github.com/jinzhu/copier"
)

type logCfg struct {
	Level string
	level logrus.Level
	File  string
}

var logConfig logCfg

type OptionalParameter struct {
	BlockChainUrl     string
	DatabaseIpAddress string
	PublicIpAddress   string
}

var OptionalConfig OptionalParameter

// Config contains all configuration read from config.yaml and network.yaml
type Config struct {
	Alias      string
	SecretKey  string
	PublicKey  string
	WalletPath string
	Channel    channelConfig
	Node       nodeConfig
	Chain      chainConfig
	// Read from the network.yaml. The key is the alias.
	Peers map[string]*netConfigEntry
}

type channelConfig struct {
	Timeout              time.Duration
	FundTimeout          time.Duration
	SettleTimeout        time.Duration
	ChallengeDurationSec uint64
}

type nodeConfig struct {
	IP              string
	Port            uint16
	DialTimeout     time.Duration
	HandleTimeout   time.Duration
	ReconnecTimeout time.Duration

	PersistencePath    string
	PersistenceEnabled bool
}

type contractSetupOption int

var contractSetupOptions [4]string = [...]string{"validate", "deploy", "validateordeploy", ""}

const (
	contractSetupOptionValidate contractSetupOption = iota
	contractSetupOptionDeploy
	contractSetupOptionValidateOrDeploy
	contractSetupOptionNotSet
)

func (option contractSetupOption) String() string {
	return contractSetupOptions[option]
}

func parseContractSetupOption(s string) (option contractSetupOption, err error) {
	for i, optionString := range contractSetupOptions {
		if s == optionString {
			option = contractSetupOption(i)
			return
		}
	}

	err = errors.New(fmt.Sprintf("Invalid value for config option 'contractsetup'. The value is '%s', but must be one of '%v'.", s, contractSetupOptions))
	log.Error().Msg(err.Error())
	return
}

type chainConfig struct {
	TxTimeout     time.Duration       // timeout duration for on-chain transactions
	ContractSetup string              // contract setup method
	contractSetup contractSetupOption //
	Adjudicator   string              // address of adjudicator contract
	adjudicator   common.Address      //
	Assetholder   string              // address of asset holder contract
	assetholder   common.Address      //
	URL           string              // URL the endpoint of your ethereum node / infura eg: ws://10.70.5.70:8546
}

type netConfigEntry struct {
	peerID      string
	peerAddress ewallet.Address
	Lel         string
	Hostname    string
	Port        uint16
}

type TsConfig struct {
	mtx    sync.Mutex
	config Config
}

var tsconfig TsConfig

// GetConfig returns a pointer to the current `Config`.
// This is needed to make viper and cobra work together.
func GetConfig() *Config {
	return &tsconfig.config
}

// SetConfig called by viper when the config file was parsed, it is alias dependend
func SetConfig() {
	if err := viper.Unmarshal(&tsconfig.config); err != nil {
		log.Error().Msg(err.Error())
	}

	var err error
	if tsconfig.config.Chain.contractSetup, err = parseContractSetupOption(tsconfig.config.Chain.ContractSetup); err != nil {
		log.Error().Msg(err.Error())
	}

	if len(tsconfig.config.Chain.Adjudicator) > 0 {
		if tsconfig.config.Chain.adjudicator, err = StrToCommonAddress(tsconfig.config.Chain.Adjudicator); err != nil {
			log.Error().Msg(err.Error())
		}
	}

	if len(tsconfig.config.Chain.Assetholder) > 0 {
		if tsconfig.config.Chain.assetholder, err = StrToCommonAddress(tsconfig.config.Chain.Assetholder); err != nil {
			log.Error().Msg(err.Error())
		}
	}

	/*
	 overwrite config from database
	*/
	mdbal.MariaDbConn.Open()
	contracts, err := mdbal.MariaDbConn.ReadContracts("default")
	if err == nil { // no error
		if contracts.Name == "default" { // data available
			if len(contracts.Adjudicator) > 0 {
				if tsconfig.config.Chain.adjudicator, err = StrToCommonAddress(contracts.Adjudicator); err != nil {
					log.Error().Msg(err.Error())
				}
			}
			if len(contracts.AssetHolder) > 0 {
				if tsconfig.config.Chain.assetholder, err = StrToCommonAddress(contracts.AssetHolder); err != nil {
					log.Error().Msg(err.Error())
				}
			}
		}
	}
	mdbal.MariaDbConn.Close()

	/*
	 changed error handling, since while starting the server
	 we don't know the user against starting demo from terminal
	*/
	for alias, peer := range tsconfig.config.Peers {
		addr, err := StrToWalletAddress(peer.peerID)
		if err != nil {
			log.Error().Msg(err.Error())
		}

		if addr != nil {
			tsconfig.config.Peers[alias].peerID = peer.peerID
		}
	}

	/*
	 check some addional  arguments from terminal and overwrite local settings by parameter from server call
	*/
	if OptionalConfig.BlockChainUrl != "" {
		log.Info().Msgf("Use block chain boot node %s", OptionalConfig.BlockChainUrl)
		tsconfig.config.Chain.URL = OptionalConfig.BlockChainUrl
	} else {
		log.Info().Msg("No alternative block chain boot node have been set, using local config")
	}

	if OptionalConfig.DatabaseIpAddress != "" {
		log.Info().Msgf("Use database node %s", OptionalConfig.DatabaseIpAddress)
	} else {
		log.Info().Msg("No database have been set, using local sessions only")
	}
}

func setConfig() {
	lvl, err := logrus.ParseLevel(logConfig.Level)
	if err != nil {
		log.Error().Msg(errors.WithMessage(err, "parsing log level").Error())
	}
	logConfig.level = lvl

	// Set the logging output file
	logger := logrus.New()
	if logConfig.File != "" {
		f, err := os.OpenFile(logConfig.File,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Error().Msg(errors.WithMessage(err, "opening logging file").Error())
		}
		logger.SetOutput(f)
	}
	logger.SetLevel(lvl)
	plog.Set(plogrus.FromLogrus(logger))
}

//Clone deep copy object details
func (c *Config) Clone() *Config {
	var cloned Config
	copier.Copy(&cloned, &c)
	peersOrg := c.Peers
	peersCp := cloned.Peers
	copier.Copy(&peersCp, &peersOrg)
	cloned.Peers = peersCp
	return &cloned
}
