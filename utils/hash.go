package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pasword string) (string, error) {
	bcrypt, err := bcrypt.GenerateFromPassword([]byte(pasword), 14)
	return string(bcrypt), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
