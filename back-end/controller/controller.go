package controller

import (
	"JUALiND/controller/router"
	"JUALiND/middleware"
	"JUALiND/repository"
	"net/http"
)

func NewMux(userRepo *repository.UserRepository, productRepo *repository.ProductRepository, orderRepo *repository.OrderRepository) *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/hallo", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}))
	m.Handle("/assets/", middleware.AllowOrigin(http.StripPrefix("/assets", http.FileServer(http.Dir("./assets")))))
	m.Handle("/api/v1/users", middleware.AllowOrigin(middleware.Method("GET", middleware.AuthMiddleware(router.GetAllUser(userRepo)))))
	m.Handle("/api/v1/user", middleware.AllowOrigin(middleware.Method("GET", middleware.AuthMiddleware(router.GetUserByID(userRepo)))))
	m.Handle("/api/v1/user/create", middleware.AllowOrigin(middleware.Method("POST", middleware.AuthMiddleware(router.CreateUser(userRepo)))))
	m.Handle("/api/v1/user/delete", middleware.AllowOrigin(middleware.Method("DELETE", middleware.AuthMiddleware(router.DeleteUserByID(userRepo)))))
	m.Handle("/api/v1/user/update", middleware.AllowOrigin(middleware.Method("PUT", middleware.AuthMiddleware(router.UpdateUser(userRepo)))))
	m.Handle("/api/v1/products", middleware.AllowOrigin(middleware.Method("GET", middleware.AuthMiddleware(router.GetAllProduct(productRepo)))))
	m.Handle("/api/v1/product", middleware.AllowOrigin(middleware.Method("GET", middleware.AuthMiddleware(router.GetProductID(productRepo)))))
	m.Handle("/api/v1/product/delete", middleware.AllowOrigin(middleware.Method("DELETE", middleware.AuthMiddleware(router.DeleteProductByID(productRepo)))))
	m.Handle("/api/v1/product/create", middleware.AllowOrigin(middleware.Method("POST", middleware.AuthMiddleware(middleware.Role("penjual", router.CreateProduct(productRepo))))))
	m.Handle("/api/v1/product/update", middleware.AllowOrigin(middleware.Method("PUT", middleware.AuthMiddleware(middleware.Role("penjual", router.UpdateProduct(productRepo))))))
	m.Handle("/api/v1/product/search", middleware.AllowOrigin(middleware.Method("GET", middleware.AuthMiddleware(router.GetProductByName(productRepo)))))
	m.Handle("/api/v1/login", middleware.AllowOrigin(middleware.Method("GET", router.LoginUser(userRepo))))
	m.Handle("/api/v1/user/token", middleware.AllowOrigin(middleware.Method("GET", middleware.AuthMiddleware(router.GetUserByToken(userRepo)))))
	m.Handle("/api/v1/user/edit/password", middleware.AllowOrigin(middleware.Method("PUT", middleware.AuthMiddleware(router.UpdateUserPassword(userRepo)))))
	m.Handle("/api/v1/user/product", middleware.AllowOrigin(middleware.Method("GET", middleware.AuthMiddleware(middleware.Role("penjual", router.GetProductByUser(productRepo))))))
	m.Handle("/api/v1/product/orders", middleware.AllowOrigin(middleware.Method("GET", middleware.AuthMiddleware(middleware.Role("penjual", router.GetAllOrdersByUser(orderRepo))))))
	return m
}
