package database

import (
	"database/sql"
	"go-embedded-database/util"

	_ "github.com/glebarez/go-sqlite"
)

var db *sql.DB
var err error

func InitDatabase() {
	db, err = sql.Open("sqlite", "datos.db")

	util.HandleError(err)
}

func CloseDatabase() {
	db.Close()
}

func PrepareStatement(query string) *sql.Stmt {
	var stmt *sql.Stmt

	stmt, err = db.Prepare(query)

	return stmt
}
