package db

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	// host =	"localhost"
	// port =	2020
	user     = "kwik"
	password = "password"
	dbname   = "doggo_db"
)

//App is the main application struct
type App struct {
	DB *sql.DB
}

//Initialize is used to open a connection to the postgres db
func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

}
