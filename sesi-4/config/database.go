package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func DB() *sql.DB {
	if db == nil {
		var err error

		dsn := getPostgresDSN()
		db, err = sql.Open("postgres", dsn)

		if nil != err {
			fmt.Println("Failed to create DB Connection ", err.Error())
		}

	}

	return db
}

func getPostgresDSN() string {
	return fmt.Sprintf(GetValue(DATABASE_CONNECTION_STRING), GetValue(DATABASE_HOST), GetValue(DATABASE_USER), GetValue(DATABASE_PASS), GetValue(DATABASE_NAME), GetValue(DATABASE_PORT), GetValue(DATABASE_SSL))
}
