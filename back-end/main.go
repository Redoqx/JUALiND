package main

import (
	"JUALiND/helper"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.db")

	if err != nil {
		panic(err)
	}

	helper.InitDB(db)
	helper.Migrate(db)
	fmt.Println("Hallo World")
}
