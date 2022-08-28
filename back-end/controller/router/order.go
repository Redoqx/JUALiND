package router

import (
	"JUALiND/helper"
	"JUALiND/models"
	"JUALiND/repository"
	"log"
	"net/http"
	"strconv"
)

func CreateOrder(orderRepo *repository.OrderRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var o models.Order
		var err error
		o.Amount, err = strconv.Atoi(r.FormValue("amount"))
		if err != nil {
			log.Println("Error on Create Order : ", err.Error())
			helper.ErrorResponseJSON(w, err, "Bad Request", http.StatusBadRequest)
			return
		}
		o.BuyerID = int(r.Context().Value("user_detail").(models.Users).ID)
		if err != nil {
			log.Println("Error on Create Order : ", err.Error())
			helper.ErrorResponseJSON(w, err, "Bad Request", http.StatusBadRequest)
			return
		}
		o.ProductID, err = strconv.Atoi(r.FormValue("id_product"))

		if err != nil {
			log.Println("Error on Create Order : ", err.Error())
			helper.ErrorResponseJSON(w, err, "Bad Request", http.StatusBadRequest)
			return
		}

		err = orderRepo.CreateOrder(o)

		if err != nil {
			log.Println("Error on Create Order : ", err.Error())
			helper.ErrorResponseJSON(w, err, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helper.SuccessResponseJSON(w, "Success Creating User", http.StatusOK)

	})
}
func GetAllOrdersByUser(orderRepo *repository.OrderRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UserData := r.Context().Value("user_detail").(models.Users)

		results, err := orderRepo.GetAllOrdersByUser(int(UserData.ID))

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Not Found", http.StatusNotFound)
			return
		}

		helper.SuccessResponseJSON(w, "Success", results)
	})
}
