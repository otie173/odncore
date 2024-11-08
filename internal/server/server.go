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
	MAX_BUFFER_SIZE int64 = 102400
)

// Opcodes for requests to client
const (
	SEND_WORLD byte = iota
	RECEIVE_WORLD

	SEND_ID
	RECEIVE_ID

	BLOCK_PACKET
	ADD_BLOCK
	REMOVE_BLOCK

	PLAYER_PACKET
	SAVE_PLAYER_DATA
	LOAD_PLAYER_DATA
)

type ServerStatus struct {
	Address          string `json:"address"`
	IdWaiting        bool   `json:"id_waiting"`
	WorldWaiting     bool   `json:"world_waiting"`
	PlayersConnected int    `json:"players_connected"`
	MaxPlayers       int    `json:"max_players"`
}

func New(address string, maxPlayersCount int) {
	m := melody.New()
	m.Config.MaxMessageSize = MAX_BUFFER_SIZE

	websocket = m
	addr = address
	maxPlayers = maxPlayersCount
}