package controller

import (
	"JUALiND/controller/router"
	"JUALiND/repository"
	"net/http"
)

func NewMux(userRepo *repository.UserRepository) *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/hallo", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}))
	m.Handle("/api/v1/users", router.GetAllUser(userRepo))
	m.Handle("/api/v1/user", router.GetUserByID(userRepo))
	m.Handle("/api/v1/user/delete", router.DeleteUserByID(userRepo))
	return m
}
