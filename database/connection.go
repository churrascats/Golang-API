package database

import (
	"API/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func CreateConnection() (conn *sql.DB, err error) {
	dbConfig := config.GetDBConfig()

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Pass, dbConfig.Database,
	)

	conn, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return
}
