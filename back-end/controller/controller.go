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
	m.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	m.Handle("/api/v1/users", router.GetAllUser(userRepo))
	m.Handle("/api/v1/user", router.GetUserByID(userRepo))
	m.Handle("/api/v1/user/create", router.CreateUser(userRepo))
	m.Handle("/api/v1/user/delete", router.DeleteUserByID(userRepo))
	m.Handle("/api/v1/user/update", router.UpdateUser(userRepo))
	m.Handle("/api/v1/products", router.GetAllProduct(productRepo))
	m.Handle("/api/v1/product", router.GetProductID(productRepo))
	m.Handle("/api/v1/product/delete", router.DeleteProductByID(productRepo))
	m.Handle("/api/v1/product/create", router.CreateProduct(productRepo))
	m.Handle("/api/v1/product/update", router.UpdateProduct(productRepo))
	m.Handle("/api/v1/product/search", router.GetProductByName(productRepo))
	return m
}
