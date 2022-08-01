package router

import (
	"JUALiND/helper"
	"JUALiND/repository"
	"net/http"
	"strconv"
)

func GetAllUser(userRepo *repository.UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users, err := userRepo.GetUsers()

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Internal Error", http.StatusInternalServerError)
			return
		}

		helper.SuccessResponseJSON(w, "Success", users)

	})
}

func GetUserByID(userRepo *repository.UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Bad Request", http.StatusBadRequest)
			return
		}

		user, err := userRepo.GetUser(id)

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Not Found", http.StatusNotFound)
			return
		}
		helper.SuccessResponseJSON(w, "Success", user)
	})
}

func DeleteUserByID(userRepo *repository.UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Bad Request", http.StatusBadRequest)
			return
		}

		err = userRepo.DeleteUser(id)

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Internal Error", http.StatusInternalServerError)
			return
		}

		helper.SuccessResponseJSON(w, "Success to delete", nil)
	})
}
