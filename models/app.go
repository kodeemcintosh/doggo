package models

import(
	"database/sql"
)

//App is the main application struct
type App struct {
	DB	*sql.DB
}