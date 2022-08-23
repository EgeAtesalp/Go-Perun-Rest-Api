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

func (db *DatabaseConnection) ReadUser(alias string) (model.User, error) {
	mDb := db.MariaDb
	user := model.User{}
	query := "SELECT * FROM users WHERE alias=?"
	stmt, err := mDb.Prepare(query)

	if err != nil {
		log.Error().Msgf("cannot prepare:%s error:%s\n", query, err)
		return model.User{}, err
	}

	err = stmt.QueryRow(alias).Scan(&user.Alias, &user.ID, &user.Secret, &user.SecretKey, &user.PrivateKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, nil
		}
		return model.User{}, err
	}

	return user, nil
}

func (db *DatabaseConnection) CreateOrUpdateUser(user model.User) error {
	mDb := db.MariaDb
	query := "REPLACE INTO users VALUES (?, ?, ?, ?, ?)"
	stmt, err := mDb.Prepare(query)

	if err != nil {
		log.Error().Msgf("cannot prepare:%s error:%s\n", query, err)
		return err
	}

	_, err = stmt.Exec(user.Alias, user.ID, user.Secret, user.SecretKey, user.PrivateKey)
	if err != nil {
		log.Error().Msgf("cannot execute:%s error:%s\n", query, err)
		return err
	}

	return nil
}

func (db *DatabaseConnection) RemoveUser(user model.User) error {
	mDb := db.MariaDb
	query := "DELETE FROM users WHERE alias=?"
	stmt, err := mDb.Prepare(query)

	if err != nil {
		log.Error().Msgf("cannot prepare:%s error:%s\n", query, err)
		return err
	}

	_, err = stmt.Exec(user.Alias)
	if err != nil {
		return err
	}

	return nil
}
