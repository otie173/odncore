package auth

import (
	"log"

	"github.com/otie173/odncore/utils/database"
)

func RegisterPlayer(nickname, password string) bool {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		log.Println(err)
		return false
	}
	database.AddPlayer(nickname, hashedPassword)
	return true
}
