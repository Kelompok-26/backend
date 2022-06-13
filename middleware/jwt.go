package middleware

import (
	"backend/constants"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(adminID int, PhoneNumber int) (string, error) {
	claims := jwt.MapClaims{}
	claims["adminID"] = adminID
	claims["PhoneNumber"] = PhoneNumber
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
