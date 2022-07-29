package main

import (
	"JUALiND/controller"
	"JUALiND/helper"
	"JUALiND/repository"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.db")

	if err != nil {
		panic(err)
	}

	helper.InitDB(db)
	helper.Migrate(db)
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	m := controller.NewMux(userRepo, productRepo)
	log.Println("Server Listening at port 8000")
	http.ListenAndServe(":8000", m)
}
