package helper

import (
	"JUALiND/models"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	User models.Users
	jwt.StandardClaims
}

var JwtKey = []byte("apacikkk")
