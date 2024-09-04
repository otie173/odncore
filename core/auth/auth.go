package auth

import "golang.org/x/crypto/bcrypt"

type Player struct {
	Nickname     string `json:"nickname"`
	PasswordHash string `json:"password_hash"`
	Salt         string `json:"salt"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
