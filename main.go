package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/glebarez/go-sqlite"
)

func main() {

	var query string
	var result int

	rand.Seed(time.Now().UnixNano())

	query = fmt.Sprintf("INSERT INTO TEST (id) VALUES (%v)", rand.Int())
	// connect
	db, err := sql.Open("sqlite", "datos.db")

	handleError(err)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY);`)

	handleError(err)

	_, err = db.Exec(query)

	handleError(err)

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
