package io

import (
	"log"

	"github.com/olahol/melody"
	"github.com/otie173/odncore/core/network/server"
)

func SetupReadHandler(s *server.Server) {
	s.Websocket.HandleMessage(func(session *melody.Session, msg []byte) {
		log.Println("Received message from", session.Request.RemoteAddr, ":", string(msg))
	})
}
