package server

import "github.com/otie173/odncore/internal/game/world"

func GetStatus() *ServerStatus {
	return &ServerStatus{
		IdWaiting:        world.IsIdWaiting,
		WorldWaiting:     world.IsWorldWaiting,
		WorldInfoWaiting: world.IsWorldInfoWaiting,
	}
}
