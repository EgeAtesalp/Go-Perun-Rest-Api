/*
 * ODS
 *
 * This is a simple REST API to access Block Chain on Ethereum and handling Smart Contracts and Payment Channel as well.
 *
 * API version: 1.0.0
 * Contact: u.kuehn@tu-berlin.de
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/rs/zerolog/log"

	pch "restapidemo/files/paymentchannel"
	"restapidemo/model"
)

//Open a paymentchannel
func Open(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	alias := readValueFromPath(r, "alias")
	var newChannel model.OpenPaymentChannel
	var result model.Result
	var e model.ModelError
	var mySess, otherSess *pch.Session

	result.Exception = e

	errIn := json.NewDecoder(r.Body).Decode(&newChannel)
	config, errCfg := setAndCheckAliasConfig(alias)
	result.Indata = newChannel

	if errCfg != nil {
		log.Error().Msgf("error read config for alias %s", alias)
		w.WriteHeader(500)
		result.Success = false
		result.Exception.ShortMessage = errCfg.Error()
		result.Exception.Error = errCfg
		e, _ := json.Marshal(result)
		w.Write(e)
		return
	}

	if errIn != nil {
		log.Info().Msg("no body available, just only do setup")
	}

	//spew.Dump(config)
	if config.Alias == "" {
		w.WriteHeader(404)
		log.Error().Str("alias-used", alias).Msg("alias unknown")
		result.Success = false
		result.Exception.ShortMessage = "alias unknown"
		e, _ := json.Marshal(result)
		w.Write(e)
		return
	}

	//fmt.Println("no error read config")
	/* check if exists any config */
	sess, err := pch.SetupAndConnect(config)
	sess.Alias = alias
	fmt.Printf("setup done for %s", alias)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(500)
		result.Success = false
		result.Exception.ShortMessage = err.Error()
		result.Exception.Error = err
		e, _ := json.Marshal(result)
		w.Write(e)
		return
	}

	if sessionExists(newChannel.Target) && newChannel.Target != "" {
		otherSess = sessions[newChannel.Target]
	}

	if otherSess == nil && newChannel.Target != "" {
		w.WriteHeader(404)
		log.Error().Msg("target is not online")
		result.Success = false
		result.Exception.ShortMessage = "target is not online"
		e, _ := json.Marshal(result)
		w.Write(e)
		return
	}

	if !sessionExists(alias) {
		sessions[alias] = &sess
	}

	// optional open channel to an already connected user
	if sessionExists(alias) && newChannel.Target != "" {
		mySess = sessions[alias]

		// connect both
		args := []string{newChannel.Target}
		sess.Backend.Connect(args, otherSess, config)
		args = []string{alias}
		otherSess.Backend.Connect(args, mySess, config)

		log.Info().Msg("connecting begins")
		_ = otherSess

		err = mySess.Backend.Open(mySess, newChannel)
		if err != nil {
			log.Error().Msgf("cannot open payment channel : %s", err)
			w.WriteHeader(500)
			result.Indata = newChannel
			result.Success = false
			result.Exception.ShortMessage = err.Error()
			result.Exception.Error = err
			e, _ := json.Marshal(result)
			w.Write(e)
			return
		}
		log.Info().Msg("connecting done")
	}

	//spew.Dump(sessions)
	//spew.Dump(result)
	b, errOut := json.Marshal(result)

	//fmt.Println("prep result")
	if errOut == nil {
		//fmt.Println("write ok")
		w.WriteHeader(http.StatusOK)
	} else {
		//fmt.Println("write 500")
		w.WriteHeader(500)
		log.Error().Msg(errOut.Error())
		e, _ := json.Marshal(errOut.Error())
		w.Write(e)
		return
	}

	//fmt.Println("header set")
	if errOut == nil {
		//fmt.Println("write body")
		w.Write(b)
	}

	log.Info().Msg("done")
}

//ValidateOrDeploy the smart contracts and store the optional new transaction hashes in
// the persitent db store
func ValidateOrDeploy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	alias := readValueFromPath(r, "alias")
	var newChannel model.OpenPaymentChannel
	var result model.Result
	var e model.ModelError
	result.Exception = e

	errIn := json.NewDecoder(r.Body).Decode(&newChannel)
	config, errCfg := setAndCheckAliasConfig(alias)
	result.Indata = newChannel

	if errCfg != nil {
		log.Error().Msgf("error read config for alias %s", alias)
		w.WriteHeader(500)
		result.Success = false
		result.Exception.ShortMessage = errCfg.Error()
		result.Exception.Error = errCfg
		e, _ := json.Marshal(result)
		w.Write(e)
		return
	}

	if errIn != nil {
		log.Info().Msg("no body available, just only do setup")
	}

	//spew.Dump(config)
	if config.Alias == "" {
		w.WriteHeader(404)
		log.Error().Str("alias-used", alias).Msg("alias unknown")
		result.Success = false
		result.Exception.ShortMessage = "alias unknown"
		e, _ := json.Marshal(result)
		w.Write(e)
		return
	}

	//fmt.Println("no error read config")
	/*
	 check if exists any config
	*/
	sess, err := pch.SetupForValidation(config)
	sess.Alias = alias
	fmt.Printf("setup done with %s", alias)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(500)
		result.Success = false
		result.Exception.ShortMessage = err.Error()
		result.Exception.Error = err
		e, _ := json.Marshal(result)
		w.Write(e)
		return
	}

	w.WriteHeader(200)
	log.Info().Msg("TBD return contracts data")

	result.Success = true
	e1, _ := json.Marshal(result)
	w.Write(e1)
}
