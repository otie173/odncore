package types

import "github.com/olahol/melody"

type ServerInterface interface {
	Start() error
	Stop() error
	SetupReadHandler()
	SendToClients(sender *melody.Session, msg []byte) error
	GetInfo() ServerInfo
}

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
