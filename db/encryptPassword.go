package db

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(pass string) (string, error) {
	cost := 8
	byte, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(byte), err
}
