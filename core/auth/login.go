package auth

import (
	"log"

	"github.com/otie173/odncore/utils/database"
)

func LoginPlayer(nickname, password string) bool {
	passwordHash, err := database.GetPasswordHash(nickname)
	if err != nil {
		log.Println(err)
	}

	validPassword := checkPassword(password, string(passwordHash))
	return validPassword
}
