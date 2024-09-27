package auth

import "github.com/otie173/odncore/utils/database"

func RegisterPlayer(nickname, password string) {
	database.AddPlayer(nickname, password)
}
