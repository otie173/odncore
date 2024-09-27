package server

func GetInfo() *ServerInfo {
	return &ServerInfo{
		Address:          addr,
		PlayersConnected: playersConnected,
		MaxPlayers:       maxPlayers,
	}
}
