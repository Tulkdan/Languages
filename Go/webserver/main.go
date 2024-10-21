package main

import (
	"database/sql"
	"example/web-server/src"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "people.db"

func main() {
	fmt.Println("Creating DB")
	file, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatalf("Error creating database %s", err)
	}

	db, err := src.CreateDB(file)
	fmt.Println("Finished creating DB")

	if err != nil {
		log.Fatalf("Error creating database %s", err)
	}

	if err := src.NewHTTPServer(":8000", db).ListenAndServe(); err != nil {
		fmt.Errorf("Error creating server %s", err)
	} else {
		fmt.Println("Listening in port 8000")
	}
}
