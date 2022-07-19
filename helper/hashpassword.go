package helper

import "golang.org/x/crypto/bcrypt"

func CreateHash(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}
