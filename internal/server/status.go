package server

import "github.com/otie173/odncore/internal/game/world"

func GetStatus() *ServerStatus {
	return &ServerStatus{
		Address:          addr,
		IdWaiting:        world.IsIdWaiting,
		WorldWaiting:     world.IsWorldWaiting,
		PlayersConnected: playersConnected,
		MaxPlayers:       maxPlayers,
	}
}
