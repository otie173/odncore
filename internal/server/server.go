package server

import (
	"github.com/olahol/melody"
)

var (
	websocket        *melody.Melody
	addr             string
	playersConnected int
	maxPlayers       int
)

const (
	maxBufferSize int64 = 102400
)

const (
	blockPacket byte = iota
	blockAdd
	blockRemove
)

type ServerStatus struct {
	IdWaiting        bool `json:"id_waiting"`
	WorldWaiting     bool `json:"world_waiting"`
	WorldInfoWaiting bool `json:"world_info_waiting"`
}

type ServerInfo struct {
	Name             string `json:"name"`
	Description      string `json:"description"`
	Address          string `json:"address"`
	PlayersConnected int    `json:"players_connected"`
	MaxPlayers       int    `json:"max_players"`
}

func InitServer(address string, maxPlayersCount int) {
	m := melody.New()
	m.Config.MaxMessageSize = maxBufferSize

	websocket = m
	addr = address
	maxPlayers = maxPlayersCount
}
