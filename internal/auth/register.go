package auth

import (
	"github.com/otie173/odncore/internal/database"
	"github.com/otie173/odncore/internal/logger"
)

func RegisterPlayer(nickname, password string) bool {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		logger.Error("Error with hash password: ", err)
		return false
	}
	database.AddPlayer(nickname, hashedPassword)
	return true
}
