package router

import (
	"JUALiND/helper"
	"JUALiND/repository"
	"net/http"
	"strconv"
)

func GetAllProduct(productRepo *repository.ProductRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		products, err := productRepo.GetProducts()

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Internal Error", http.StatusInternalServerError)
			return
		}

		helper.SuccessResponseJSON(w, "Success", products)
	})
}

func GetProductID(productRepo *repository.ProductRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Bad Request", http.StatusBadRequest)
			return
		}

		product, err := productRepo.GetProduct(id)

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Not Found", http.StatusNotFound)
			return
		}

		helper.SuccessResponseJSON(w, "Success", product)
	})
}

func DeleteProductByID(productRepo *repository.ProductRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Bad Request", http.StatusBadRequest)
			return
		}

		err = productRepo.DeleteProduct(id)

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Internal Error", http.StatusInternalServerError)
			return
		}

		helper.SuccessResponseJSON(w, "Success", nil)

	})
}
