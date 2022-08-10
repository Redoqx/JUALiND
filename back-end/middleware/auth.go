package middleware

import (
	"JUALiND/helper"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			log.Println(authorizationHeader)
			helper.ErrorResponseJSON(w, fmt.Errorf("error : bad request"), "Token not found", http.StatusBadRequest)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		userClaims := helper.UserClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &userClaims, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("signing method invalid")
			}

			return helper.JwtKey, nil
		})
		if err != nil {
			log.Println("test")
			helper.ErrorResponseJSON(w, err, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			helper.ErrorResponseJSON(w, fmt.Errorf("token is invalid"), "Unauthorized", http.StatusUnauthorized)
			return
		}
		userData := userClaims.User
		fmt.Println(userData)
		ctx := context.WithValue(context.Background(), "user_detail", userData)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
