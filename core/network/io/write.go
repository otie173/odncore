package io

import (
	"github.com/olahol/melody"
	"github.com/otie173/odncore/core/network/types"
)

func SendToClients(s *types.Server, sender *melody.Session, msg []byte) error {
	return s.Websocket.BroadcastFilter(msg, func(session *melody.Session) bool {
		return sender != session
	})
}
