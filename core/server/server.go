package server

import (
	"github.com/olahol/melody"
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
	return &Server{
		Websocket:  melody.New(),
		Addr:       addr,
		MaxPlayers: maxPlayers,
	}
}
