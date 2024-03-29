package main

import (
	"JUALiND/controller"
	"JUALiND/helper"
	"JUALiND/repository"
	"database/sql"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	rand.Seed(time.Now().Unix())
	db, err := sql.Open("sqlite3", "file:./database.db?_foreign_keys=true")

	if err != nil {
		panic(err)
	}

	PORT := os.Getenv("PORT")

	if len(PORT) < 1 {
		PORT = "8000"
	}

	helper.InitDB(db)
	helper.Migrate(db)
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	m := controller.NewMux(userRepo, productRepo, orderRepo)
	log.Println("Server Listening at port 8000")
	http.ListenAndServe(":"+PORT, m)
}
