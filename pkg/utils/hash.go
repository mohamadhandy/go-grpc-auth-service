package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	btyes, _ := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(btyes)
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
