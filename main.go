package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/glebarez/go-sqlite"
)

func main() {

	var result int
	// connect
	db, err := sql.Open("sqlite", "datos.db")

	handleError(err)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY);`)

	handleError(err)

	//_, err = db.Exec(`INSERT INTO test (id) VALUES(1);`)

	handleError(err)

	// get SQLite version
	rows, err := db.Query("SELECT id FROM test;")

	handleError(err)

	defer rows.Close()

	rows.Next()

	err = rows.Scan(&result)

	handleError(err)

	fmt.Println(result)

	db.Close()
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
