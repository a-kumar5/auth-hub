package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePassword(hashedPassword, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err
}
