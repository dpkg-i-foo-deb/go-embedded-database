package database

import (
	"database/sql"
	"go-embedded-database/util"
)

var createTestTableStmt *sql.Stmt
var insertCatStmt *sql.Stmt

func initStatements() {
	createTestTableStmt = PrepareStatement(`CREATE TABLE 
											IF NOT EXISTS cat 
												(id INTEGER PRIMARY KEY AUTOINCREMENT,
												name VARCHAR(80) NOT NULL)`)
	insertCatStmt = PrepareStatement(`INSERT INTO cat (name) 
										VALUES (?)`)

	createTables()
}

func InsertCat(name string) {
	_, err := insertCatStmt.Exec(name)

	util.HandleError(err)
}

func createTables() {
	_, err := createTestTableStmt.Exec()

	util.HandleError(err)
}
