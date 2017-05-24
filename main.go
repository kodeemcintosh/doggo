package main

import (
	// "database/sql"
	// "flag"
	// "fmt"
	db "github.com/kvmac/doggo/db"
	_ "github.com/lib/pq"
	// "os"
)

const (
	username = "doggo"
	password = "pupper"
	dbname   = "doggo_db"
)

func main() {

	a := db.App{}
	// os.Setenv()

	a.Initialize(username, password, dbname)

	// Args := os.Args

}
