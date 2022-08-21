package router

import (
	"JUALiND/helper"
	"JUALiND/models"
	"JUALiND/repository"
	"net/http"
)

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
