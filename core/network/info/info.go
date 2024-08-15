package info

import "github.com/otie173/odncore/core/network/types"

func GetInfo(s *types.Server) types.ServerInfo {
	return types.ServerInfo{
		Address:          s.Addr,
		PlayersConnected: s.PlayersConnected,
		MaxPlayers:       s.MaxPlayers,
	}
}
