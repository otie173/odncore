package io

import (
	"github.com/olahol/melody"
	"github.com/otie173/odncore/core/network/server"
)

func SendToClients(s *server.Server, sender *melody.Session, msg []byte) error {
	return s.Websocket.BroadcastFilter(msg, func(session *melody.Session) bool {
		return sender != session
	})
}
