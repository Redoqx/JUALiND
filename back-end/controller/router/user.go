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
	"time"

	"github.com/golang-jwt/jwt"
)

type tokenBody struct {
	Token string `json:"token"`
}

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
func GetUserByToken(userRepo *repository.UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UserData := r.Context().Value("user_detail").(models.Users)

		expTime := time.Now().Add(60 * time.Minute)
		newTokenString, err := generateUserToken(UserData, expTime)

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helper.SuccessResponseJSON(w, "Success", map[string]interface{}{
			"user":  UserData,
			"token": newTokenString,
		})
		return
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

		user, err := userRepo.GetUser(id)

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Internal Error", http.StatusInternalServerError)
			return
		}

		if user.ImageLoc.Valid {
			os.Remove(user.ImageLoc.String)
		}
		err = userRepo.DeleteUser(id)

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Internal Error", http.StatusInternalServerError)
			return
		}

		helper.SuccessResponseJSON(w, "Success to delete", nil)
	})
}
func generateUserToken(user models.Users, expTime time.Time) (string, error) {
	claims := &helper.UserClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(helper.JwtKey)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil

}
func LoginUser(userRepo *repository.UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := userRepo.GetUserByEmail(email)
		if err != nil {
			helper.ErrorResponseJSON(w, err, "User Not Found", http.StatusNotFound)
			return
		}

		if !helper.CheckPasswordHash(password, user.Password) {
			helper.ErrorResponseJSON(w, fmt.Errorf("wrong password"), "Wrong Password", http.StatusUnauthorized)
			return
		}

		expTime := time.Now().Add(30 * time.Minute)
		tokenString, err := generateUserToken(*user, expTime)

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		helper.SuccessResponseJSON(w, "login Success", tokenBody{Token: tokenString})
	})
}

func CreateUser(userRepo *repository.UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		password := r.FormValue("password")
		email := r.FormValue("email")
		imageFile, imageHeader, err := r.FormFile("image")

		if len(name) < 1 ||
			len(password) < 1 ||
			len(email) < 1 {
			helper.ErrorResponseJSON(w, fmt.Errorf("error : request invalid"), "some field are not filled", http.StatusBadRequest)
			return
		}
		password, _ = helper.HashPassword(password)
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
			if !helper.ImageIsJpgOrPng(imageHeader) {
				helper.ErrorResponseJSON(w, fmt.Errorf("error : request invalid"), "image must be in png or jpg format", http.StatusBadRequest)
				return
			}
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
			name := r.FormValue("name")
			password := r.FormValue("password")
			email := r.FormValue("email")
			imageFile, imageHeader, err := r.FormFile("image")

			if len(name) < 1 ||
				len(password) < 1 ||
				len(email) < 1 {
				helper.ErrorResponseJSON(w, fmt.Errorf("error : request invalid"), "some field are not filled", http.StatusBadRequest)
				return
			}
			password, _ = helper.HashPassword(password)
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
				if !helper.ImageIsJpgOrPng(imageHeader) {
					helper.ErrorResponseJSON(w, fmt.Errorf("error : request invalid"), "image must be in png or jpg format", http.StatusBadRequest)
					return
				}
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
		}
	})
}
func UpdateUserPassword(userRepo *repository.UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		password := r.FormValue("password")

		if len(password) < 1 {
			helper.ErrorResponseJSON(w, fmt.Errorf("error : request invalid"), "some field are not filled", http.StatusBadRequest)
			return
		}

		userData := r.Context().Value("user_detail").(models.Users)
		hashedPass, err := helper.HashPassword(password)

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Internal Server Error", http.StatusInternalServerError)
			return

		}

		err = userRepo.UpdateUserPassword(int(userData.ID), hashedPass)
		if err != nil {
			helper.ErrorResponseJSON(w, err, "Internal Server Error", http.StatusInternalServerError)
			return

		}

		userData.Password = hashedPass
		expTime := time.Now().Add(30 * time.Minute)
		newToken, err := generateUserToken(userData, expTime)

		if err != nil {
			helper.ErrorResponseJSON(w, err, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helper.SuccessResponseJSON(w, "Success", tokenBody{
			Token: newToken,
		})

	})
}
func UpdateUser(userRepo *repository.UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id, err_id := strconv.Atoi(r.FormValue("id"))
		name := r.FormValue("name")
		email := r.FormValue("email")
		imageFile, imageHeader, err := r.FormFile("image")

		if len(name) < 1 ||
			err_id != nil ||
			id < 1 ||
			len(email) < 1 {
			helper.ErrorResponseJSON(w, fmt.Errorf("error : request invalid"), "some field are not filled", http.StatusBadRequest)
			return
		}

		currentUser, _ := userRepo.GetUser(id)
		var u models.Users
		if err != nil {
			// Langsung save user tanpa gambar
			u.ID = uint(id)
			u.Name = name
			u.Email = email
			u.ImageLoc = helper.StringToNullString(currentUser.ImageLoc.String)
			err = userRepo.UpdateUser(u)
			if err != nil {
				helper.ErrorResponseJSON(w, err, "Internal Server Error", http.StatusInternalServerError)
			}
			helper.SuccessResponseJSON(w, "success", nil)
		} else {
			// save user serta gambarnya
			if !helper.ImageIsJpgOrPng(imageHeader) {
				helper.ErrorResponseJSON(w, fmt.Errorf("error : request invalid"), "image must be in png or jpg format", http.StatusBadRequest)
				return
			}
			u.ID = uint(id)
			u.Name = name
			u.Email = email
			fileLocation := helper.UploadImage(imageFile, imageHeader)
			u.ImageLoc = helper.StringToNullString(r.Host + "/" + fileLocation)

			if currentUser.ImageLoc.Valid {
				helper.RemoveFile(r, currentUser.ImageLoc.String)
			}

			log.Println(r.Host)
			err = userRepo.UpdateUser(u)
			if err != nil {
				helper.ErrorResponseJSON(w, err, "Internal Server Error", http.StatusInternalServerError)
				os.Remove(fileLocation)
				return
			}
			helper.SuccessResponseJSON(w, "success", nil)
		}
	})
}
