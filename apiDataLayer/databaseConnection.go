package apiDataLayer

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	HOST = "localhost"
	PORT = "5432"
	USER = "jarrettnester"
	PASSWORD = "janejohn123"
	DBNAME = "susannester"
)

var DB *sql.DB

func RetrieveDatabase() (db *sql.DB, err error) {

	if DB == nil {
		err = setDatabase()
	}

	return DB, err

}

func setDatabase() (err error) {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DBNAME)

	DB, err = sql.Open("postgres", psqlInfo)

	return
}
