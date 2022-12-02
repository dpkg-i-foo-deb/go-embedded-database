package main

import (
	"go-embedded-database/database"
	"go-embedded-database/util"
	"os"
)

func main() {
	database.InitDatabase()

	database.InitStatements()

	createTable()

	database.CloseDatabase()

	os.Exit(0)
}

func createTable() {
	_, err := database.CreateTestTableStmt.Exec()

	util.HandleError(err)
}
