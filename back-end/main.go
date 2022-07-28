package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	_, err := sql.Open("sqlite3", "./database.db")

	if err != nil {
		panic(err)
	}

	fmt.Println("Hallo World")
}
