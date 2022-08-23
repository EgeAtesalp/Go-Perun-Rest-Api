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
	"container/list"
	sql "database/sql"
	"encoding/json"
	"restapidemo/model"

	"github.com/rs/zerolog/log"

	_ "github.com/go-sql-driver/mysql"
)

func (db *DatabaseConnection) ReadSDPEntriesBySender(alias string) (*list.List, error) {
	mDb := db.MariaDb
	query := "SELECT * FROM simple_data_protocol WHERE sender_alias=? AND msg_id_ref=0"
	stmt, err := mDb.Prepare(query)

	if err != nil {
		log.Error().Msgf("cannot prepare:%s error:%s\n", query, err)
		return nil, err
	}

	data := list.New()
	rows, err := stmt.Query(alias)

	if err != nil {
		if err == sql.ErrNoRows {
			return data, nil
		}
		return data, err
	}

	for rows.Next() {
		row := model.SimpleDataProtocol{}
		err := rows.Scan(&row.ID, &row.SenderAlias, &row.ReceiverAlias, &row.MsgId, &row.PaymentChannel, &row.MsgIdRef, &row.Data, &row.LastUpdated)

		if err != nil {
			log.Error().Msgf("error occurred while reading data %s\n", err)
		}

		data.PushBack(row)
	}

	rows.Close()
	return data, nil
}

func (db *DatabaseConnection) CreateSDPEntry(sdp model.SimpleDataProtocol) error {
	mDb := db.MariaDb
	query := `INSERT INTO simple_data_protocol 
				(sender_alias, receiver_alias, msg_id, paymentchannel, msg_id_ref, data)
				VALUES (?, ?, ?, ?, ?, ?)`
	stmt, err := mDb.Prepare(query)

	if err != nil {
		log.Error().Msgf("cannot prepare:%s error:%s\n", query, err)
		return err
	}

	data, _ := json.Marshal(sdp.Data)

	res, err := stmt.Exec(sdp.SenderAlias, sdp.ReceiverAlias, sdp.MsgId, sdp.PaymentChannel, sdp.MsgIdRef, data)
	if err != nil {
		log.Error().Msgf("cannot prepare:%s error:%s\n", query, err)
		return err
	}

	sdp.ID, err = res.LastInsertId()
	if err != nil {
		log.Error().Msgf("error while reading  new id:%d error:%s\n", sdp.ID, err)
		return err
	}

	return nil
}

func (db *DatabaseConnection) UpdateSDPEntry(sdp model.SimpleDataProtocol) error {
	mDb := db.MariaDb
	query := `UPDATE simple_data_protocol 
				SET sender_alias=?, receiver_alias=?, msg_id=?, paymentchannel=?, msg_id_ref=?, data=?
				WHERE ID=?`
	stmt, err := mDb.Prepare(query)

	if err != nil {
		log.Error().Msgf("cannot prepare:%s error:%s\n", query, err)
		return err
	}

	data, _ := json.Marshal(sdp.Data)

	_, err = stmt.Exec(sdp.SenderAlias, sdp.ReceiverAlias, sdp.MsgId, sdp.MsgIdRef, sdp.PaymentChannel, data, sdp.ID)
	if err != nil {
		log.Error().Msgf("cannot prepare:%s error:%s\n", query, err)
		return err
	}

	return nil
}

func (db *DatabaseConnection) RemoveSDPEntry(sdp model.SimpleDataProtocol) error {
	mDb := db.MariaDb
	query := "DELETE FROM simple_data_protocol WHERE ID=?"
	stmt, err := mDb.Prepare(query)

	if err != nil {
		log.Error().Msgf("cannot prepare:%s error:%s\n", query, err)
		return err
	}

	_, err = stmt.Exec(sdp.ID)
	if err != nil {
		log.Error().Msgf("cannot prepare:%s error:%s\n", query, err)
		return err
	}

	return nil
}
