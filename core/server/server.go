package server

import (
	"github.com/olahol/melody"
)

// Opcodes for requests to client
const (
	SEND_WORLD byte = iota
)

type ServerInfo struct {
	Address          string `json:"addres"`
	PlayersConnected int    `json:"players_connected"`
	MaxPlayers       int    `json:"max_players"`
}

type Server struct {
	Websocket        *melody.Melody
	Addr             string
	PlayersConnected int
	MaxPlayers       int
}

func New(addr string, maxPlayers int) *Server {
	m := melody.New()
	m.Config.MaxMessageSize = 102400

	return &Server{
		Websocket:  m,
		Addr:       addr,
		MaxPlayers: maxPlayers,
	}
}
