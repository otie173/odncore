package server

func (s *Server) GetInfo() *ServerInfo {
	return &ServerInfo{
		Address:          s.Addr,
		PlayersConnected: s.PlayersConnected,
		MaxPlayers:       s.MaxPlayers,
	}
}
