/*
 * ODS
 *
 * This is a simple REST API to access Block Chain on Ethereum and handling Smart Contracts and Payment Channel as well.
 * mdbal provides access to a mariadb, to handle:
 * 	- simple data exchange protocol
 *	- session handling
 * 	- configuration handling
 *
 * API version: 1.0.0
 * Contact: u.kuehn@tu-berlin.de
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package mdbal

import (
	sql "database/sql"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"

	_ "github.com/go-sql-driver/mysql"
)

// OdsApp is a payment app.
type DatabaseConnection struct {
	Addr     string
	Port     int64
	User     string
	Password string
	Database string

	Error   error
	MariaDb *sql.DB
}

var MariaDbConn = DatabaseConnection{Addr: "127.0.0.1", Port: 33306, User: "rest-api", Password: "rest-api", Database: "go-perun-rest-api"}

//Open connect to the configured database
func (db *DatabaseConnection) Open() {
	connString := db.User + ":" + db.Password + "@tcp(" + db.Addr + ":" + strconv.FormatInt(db.Port, 10) + ")/" + db.Database + "?parseTime=true"
	dbConn, err := sql.Open("mysql", connString)

	if err != nil {
		log.Error().Msgf("Could not open db connection:%s \n using conn: %s \n", err, connString)
		db.Error = err
		return
	}

	dbConn.SetConnMaxLifetime(time.Minute * 3)
	dbConn.SetMaxOpenConns(30)
	dbConn.SetMaxIdleConns(30)

	db.MariaDb = dbConn
}

func (db *DatabaseConnection) Close() {
	db.MariaDb.Close()
}