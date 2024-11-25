package auth

import (
	"github.com/otie173/odncore/internal/utils/database"
	"github.com/otie173/odncore/internal/utils/logger"
)

func RegisterPlayer(nickname, password string) bool {
	hashedPassword, err := GenerateHash(password)
	if err != nil {
		logger.Error("Error with hash password: ", err)
		return false
	}
	database.AddPlayer(nickname, string(hashedPassword))
	return true
}
