package router

import (
	"JUALiND/helper"
	"JUALiND/models"
	"JUALiND/repository"
	"fmt"
	"log"
	"net/http"
	"os"
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

func CreateUser(userRepo *repository.UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		password := r.FormValue("password")
		email := r.FormValue("email")
		imageFile, imageHeader, err := r.FormFile("image")

		if len(name) < 1 &&
			len(password) < 1 &&
			len(email) < 1 {
			helper.ErrorResponseJSON(w, fmt.Errorf("Error : request invalid"), "some field are not filled", http.StatusBadRequest)
			return
		}
		var u models.Users
		if err != nil {
			// Langsung save user tanpa gambar
			u.Name = name
			u.Email = email
			u.Password = password
			u.ImageLoc = helper.StringToNullString("")
			err = userRepo.CreateUser(u)
			if err != nil {
				helper.ErrorResponseJSON(w, err, "Internal Server Error", http.StatusInternalServerError)
			}
			helper.SuccessResponseJSON(w, "success", nil)
		} else {
			// save user serta gambarnya
			u.Name = name
			u.Email = email
			u.Password = password
			fileLocation := helper.UploadImage(imageFile, imageHeader)
			u.ImageLoc = helper.StringToNullString(r.Host + "/" + fileLocation)
			log.Println(r.Host)
			err = userRepo.CreateUser(u)
			if err != nil {
				helper.ErrorResponseJSON(w, err, "Internal Server Error", http.StatusInternalServerError)
				os.Remove(fileLocation)
				return
			}
			helper.SuccessResponseJSON(w, "success", nil)
		}
		return
	})
}
