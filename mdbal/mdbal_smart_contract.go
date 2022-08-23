/*
 * ODS
 *
 * This is a simple REST API to access Block Chain on Ethereum and handling Smart Contracts and Payment Channel as well.
 * mdbal provides access to a mariadb, to handle:
 *	- user handling
 *
 * API version: 1.0.0
 * Contact: u.kuehn@tu-berlin.de
 */

package mdbal

import (
	sql "database/sql"
	"restapidemo/model"

	"github.com/rs/zerolog/log"

	_ "github.com/go-sql-driver/mysql"
)

func (db *DatabaseConnection) ReadContracts(name string) (model.ContractsPairPaymentchannel, error) {
	mDb := db.MariaDb
	contracts := model.ContractsPairPaymentchannel{}
	query := "SELECT * FROM smart_contract_hashes WHERE name=?"
	stmt, err := mDb.Prepare(query)

	if err != nil {
		log.Error().Msgf("cannot prepare:%s error:%s\n", query, err)
		return model.ContractsPairPaymentchannel{}, err
	}

	err = stmt.QueryRow(name).Scan(&contracts.Name, &contracts.AssetHolder, &contracts.Adjudicator)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.ContractsPairPaymentchannel{}, nil
		}
		return model.ContractsPairPaymentchannel{}, err
	}

	return contracts, nil
}

func (db *DatabaseConnection) CreateOrUpdateContracts(contracts model.ContractsPairPaymentchannel) error {
	mDb := db.MariaDb
	query := "REPLACE INTO smart_contract_hashes VALUES (?, ?, ?)"
	stmt, err := mDb.Prepare(query)

	if err != nil {
		log.Error().Msgf("cannot prepare:%s error:%s\n", query, err)
		return err
	}

	_, err = stmt.Exec(contracts.Name, contracts.AssetHolder, contracts.Adjudicator)
	if err != nil {
		return err
	}

	return nil
}

func (db *DatabaseConnection) RemoveContracts(contracts model.ContractsPairPaymentchannel) error {
	mDb := db.MariaDb
	query := "DELETE FROM smart_contract_hashes WHERE name=?"
	stmt, err := mDb.Prepare(query)

	if err != nil {
		log.Error().Msgf("cannot prepare:%s error:%s\n", query, err)
		return err
	}

	_, err = stmt.Exec(contracts.Name)
	if err != nil {
		return err
	}

	return nil
}
