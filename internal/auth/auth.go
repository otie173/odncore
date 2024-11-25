package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	HASH_COST int = 14
)

type PlayerAuth struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func GenerateHash(password string) ([]byte, error) {
	if len(password) == 0 {
		return nil, errors.New("Password cant be empty")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), HASH_COST)
	return hash, err
}

func CheckPassword(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
