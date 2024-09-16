package db

import (
	"backend/utils"
	"database/sql"
	_ "embed"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var MyDB *sql.DB = nil

//go:embed init.sql
var initSQL string

func Connect() {
	log.Printf("Connecting to the database")

	// time.Sleep(10 * time.Second)

	username, ok1 := os.LookupEnv("POSTGRES_USERNAME")
	password, ok2 := os.LookupEnv("POSTGRES_PASSWORD")
	databaseIP, ok3 := os.LookupEnv("POSTGRES_HOST")
	databasePort, ok4 := os.LookupEnv("POSTGRES_PORT")
	databaseName, ok5 := os.LookupEnv("POSTGRES_DATABASE")

	if !ok1 || !ok2 || !ok3 || !ok4 || !ok5 {
		log.Printf("Database environment variables are not set")
		return
	}

	log.Printf(username)
	log.Printf(password)
	log.Printf(databaseIP)
	log.Printf(databaseName)

	connStr := "user=" + username + " password=" + password + " dbname=" + databaseName + " host=" + databaseIP + " port=" + databasePort + " sslmode=disable"

	log.Printf(connStr)

	var err error

	MyDB, err = sql.Open("postgres", connStr)

	utils.ErrorHandler(err, "Error while connecting to the database")

	setupDB()
}

func setupDB() {
	if MyDB == nil {
		log.Printf("Database is not connected")
		return
	}

	_, err := MyDB.Exec(initSQL)

	log.Println("[INFO]: Creating the table")

	utils.ErrorHandler(err, "Error while creating the table")
}
