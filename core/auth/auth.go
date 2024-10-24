package auth

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	HASH_COST int = 14
)

type PlayerAuth struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), HASH_COST)
	return string(bytes), err
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
