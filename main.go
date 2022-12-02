package main

import (
	"go-embedded-database/database"
	"os"
)

func main() {
	database.InitDatabase()

	database.InitStatements()

	database.InsertCat("Michi")

	database.CloseDatabase()

	os.Exit(0)
}
