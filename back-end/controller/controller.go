package controller

import (
	"JUALiND/controller/router"
	"JUALiND/middleware"
	"JUALiND/repository"
	"net/http"
)

func NewMux(userRepo *repository.UserRepository, productRepo *repository.ProductRepository) *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/hallo", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}))
	m.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	m.Handle("/api/v1/users", middleware.Method("GET", middleware.AuthMiddleware(router.GetAllUser(userRepo))))
	m.Handle("/api/v1/user", middleware.Method("GET", middleware.AuthMiddleware(router.GetUserByID(userRepo))))
	m.Handle("/api/v1/user/create", middleware.Method("POST", middleware.AuthMiddleware(router.CreateUser(userRepo))))
	m.Handle("/api/v1/user/delete", middleware.Method("DELETE", middleware.AuthMiddleware(router.DeleteUserByID(userRepo))))
	m.Handle("/api/v1/user/update", middleware.Method("PUT", middleware.AuthMiddleware(router.UpdateUser(userRepo))))
	m.Handle("/api/v1/products", middleware.Method("GET", middleware.AuthMiddleware(router.GetAllProduct(productRepo))))
	m.Handle("/api/v1/product", middleware.Method("GET", middleware.AuthMiddleware(router.GetProductID(productRepo))))
	m.Handle("/api/v1/product/delete", middleware.Method("DELETE", middleware.AuthMiddleware(router.DeleteProductByID(productRepo))))
	m.Handle("/api/v1/product/create", middleware.Method("POST", middleware.AuthMiddleware(router.CreateProduct(productRepo))))
	m.Handle("/api/v1/product/update", middleware.Method("PUT", middleware.AuthMiddleware(router.UpdateProduct(productRepo))))
	m.Handle("/api/v1/product/search", middleware.Method("GET", middleware.AuthMiddleware(router.GetProductByName(productRepo))))
	m.Handle("/api/v1/login", router.LoginUser(userRepo))
	m.Handle("/api/v1/user/token", middleware.Method("GET", middleware.AuthMiddleware(router.GetUserByToken(userRepo))))
	m.Handle("/api/v1/user/edit/password", middleware.Method("PUT", middleware.AuthMiddleware(router.UpdateUserPassword(userRepo))))
	return m
}
