package info

import (
	"github.com/otie173/odncore/core/network/server"
)

func GetInfo(s *server.Server) *server.ServerInfo {
	return &server.ServerInfo{
		Address:          s.Addr,
		PlayersConnected: s.PlayersConnected,
		MaxPlayers:       s.MaxPlayers,
	}
}
