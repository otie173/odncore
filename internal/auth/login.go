package auth

import (
	"github.com/otie173/odncore/internal/utils/database"
	"github.com/otie173/odncore/internal/utils/logger"
)

func LoginPlayer(nickname, password string) bool {
	passwordHash, err := database.GetPasswordHash(nickname)
	if err != nil {
		logger.Error("Error with get password hash: ", err)
	}

	validPassword := checkPassword(password, string(passwordHash))
	return validPassword
}
