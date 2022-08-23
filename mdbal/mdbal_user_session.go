/*
 * ODS
 *
 * This is a simple REST API to access Block Chain on Ethereum and handling Smart Contracts and Payment Channel as well.
 * mdbal provides access to a mariadb, to handle:
 *	- session handling
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

func (db *DatabaseConnection) ReadUserSession(alias string) (model.UserSession, error) {
	mDb := db.MariaDb
	sess := model.UserSession{}
	query := "SELECT * FROM user_session WHERE alias=?"
	stmt, err := mDb.Prepare(query)

	if err != nil {
		log.Error().Msgf("cannot prepare:%s error:%s\n", query, err)
		return model.UserSession{}, err
	}

	err = stmt.QueryRow(alias).Scan(&sess.Alias, &sess.Ip_Addr, &sess.Port)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.UserSession{}, nil
		}
		return model.UserSession{}, err
	}

	return sess, nil
}

func (db *DatabaseConnection) CreateOrUpdateUserSession(sess model.UserSession) error {
	mDb := db.MariaDb
	query := "REPLACE INTO user_session VALUES (?, ?, ?)"
	stmt, err := mDb.Prepare(query)

	if err != nil {
		log.Error().Msgf("cannot prepare:%s error:%s\n", query, err)
		return err
	}

	_, err = stmt.Exec(sess.Alias, sess.Ip_Addr, sess.Port)
	if err != nil {
		return err
	}

	return nil
}

func (db *DatabaseConnection) RemoveUserSession(sess model.UserSession) error {
	mDb := db.MariaDb
	query := "DELETE FROM user_session WHERE alias=?"
	stmt, err := mDb.Prepare(query)

	if err != nil {
		log.Error().Msgf("cannot prepare:%s error:%s\n", query, err)
		return err
	}

	_, err = stmt.Exec(sess.Alias)
	if err != nil {
		return err
	}

	return nil
}
