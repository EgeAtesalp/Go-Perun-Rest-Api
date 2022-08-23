package paymentchannel

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:              "ods",
	Short:            "ODS umbrella executable",
	Long:             "Umbrella project for demonstrators and tests of the Perun Project.",
	PersistentPreRun: runRoot,
}

var cfgFile, cfgNetFile, aliasCfgFile string

//SetConfigFile does
func SetConfigFile(file string) {
	aliasCfgFile = file
	initConfig()
}

//ExecuteReadConfig does
func ExecuteReadConfig() {
	initConfig()
}

func runRoot(c *cobra.Command, args []string) {
	setConfig()
}

//init init config
func init() {
	log.Info().Msg("run init")
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "General config file")
	rootCmd.PersistentFlags().StringVar(&aliasCfgFile, "alias", "", "Individual config file")
	rootCmd.PersistentFlags().StringVar(&cfgNetFile, "network", "network.yaml", "Network config file")
	rootCmd.PersistentFlags().StringVar(&logConfig.Level, "log-level", "all", "Logrus level")
	viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))
	rootCmd.PersistentFlags().StringVar(&logConfig.File, "log-file", "", "log file")
	viper.BindPFlag("log.file", rootCmd.PersistentFlags().Lookup("log-file"))
}

// initConfig reads the config and sets the loglevel.
// The configuration will be parsed in each API call,
// because each user requires other data
//TODO: this solution holds for on-time usage, only
func initConfig() {
	tsconfig.mtx.Lock()
	log.Info().Msg("run init config")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Info().Msgf("current path :", dir)
	
	// Load config files
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Warn().Msgf("Error reading config file, %s", err)
	}

	viper.SetConfigFile(aliasCfgFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Warn().Msgf("Error reading config file, %s", err)
	}

	viper.SetConfigFile(cfgNetFile)
	if err := viper.MergeInConfig(); err != nil {
		log.Fatal().Msgf("Error reading network config file, %s", err)
	}

	SetConfig()
	tsconfig.mtx.Unlock()
}
