package utils

import(
	"database/sql"
	"log"
)

func ErrCheck(err error) {
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("No Results Found")
		} else {
			log.Println("Oh shit you've got an error, Dawg!")
			panic(err)
		}
	}
}