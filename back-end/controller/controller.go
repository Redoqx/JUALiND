package controller

import (
	"JUALiND/controller/router"
	"JUALiND/repository"
	"net/http"
)

func NewMux(userRepo *repository.UserRepository, productRepo *repository.ProductRepository) *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/hallo", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}))
	m.Handle("/api/v1/users", router.GetAllUser(userRepo))
	m.Handle("/api/v1/user", router.GetUserByID(userRepo))
	m.Handle("/api/v1/user/delete", router.DeleteUserByID(userRepo))
	m.Handle("/api/v1/products", router.GetAllProduct(productRepo))
	m.Handle("/api/v1/product", router.GetProductID(productRepo))
	m.Handle("/api/v1/product/delete", router.DeleteProductByID(productRepo))
	return m
}
