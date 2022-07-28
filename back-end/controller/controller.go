package controller

import (
	"database/sql"
	"net/http"
)

type Controller struct {
	mux *http.ServeMux
}

func NewController(db *sql.DB) *Controller {
	return &Controller{}
}
