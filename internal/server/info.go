package server

import "github.com/otie173/odncore/internal/utils/config"

func GetInfo(cfg config.Config) *ServerInfo {
	return &ServerInfo{
		Name:             cfg.ServerName,
		Description:      cfg.ServerDescription,
		Address:          cfg.Address,
		PlayersConnected: playersConnected,
		MaxPlayers:       maxPlayers,
	}
}
