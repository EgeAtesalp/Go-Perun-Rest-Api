package swagger

import (
	e "errors"
	"net/http"
	
	"github.com/rs/zerolog/log"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"

	pch "restapidemo/files/paymentchannel"

	
)

func readValueFromPath(r *http.Request, key string) (value string) {
	//read alias
	params := mux.Vars(r)
	value = params[key]
	log.Info().Msg(key + " : " + value)
	return value
}

func setAndCheckAliasConfig(alias string) (config *pch.Config, err error) {
	pch.SetConfigFile(alias + ".yaml")
	err = nil
	config = pch.GetConfig()
	if alias != config.Alias {
		err = e.New("Could not load alias config file")
		//If not found locally, check redis database for user session
		
		//
	}

	if err != nil {
		log.Info().Str("config", spew.Sdump(config)).Msg("current config")
	}

	return config, err
}
