package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/TwiN/go-color"
	"github.com/gookit/config/v2"
	_ "github.com/lib/pq"
	db "rlesjak.com/ha-scheduler/database/generated"
)

var Q *db.Queries
var DB *sql.DB

func ConnectDatabase(config config.Config) {

	// Check if DB_HOST env variable exists to determine if
	// required env variables are available

	// Assemble the Data Source Name string
	var dataSourceName string
	password := config.String("env_db_pwd", "")
	if password == "" {
		dataSourceName = fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable",
			config.String("env_db_host"),
			config.String("env_db_user"),
			config.String("env_db_name"),
			config.String("env_db_port"),
		)
	} else {
		dataSourceName = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			config.String("env_db_host"),
			config.String("env_db_user"),
			password,
			config.String("env_db_name"),
			config.String("env_db_port"),
		)
	}

	// Display datasource string in console
	// security risk, remove. ...
	log.Println(color.InWhiteOverBlue("Datasource: " + dataSourceName))

	database, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(color.InWhiteOverRed("ERROR CONNECTING TO DATABASE \n"), err)
		return
	}

	if err := database.Ping(); err != nil {
		log.Fatal(color.InWhiteOverRed("ERROR PINGING DB \n"), err)
		return
	}

	Q = db.New(database)
	DB = database
}
