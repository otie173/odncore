package server

import (
	"log"

	"github.com/olahol/melody"
	"github.com/otie173/odncore/core/game/world"
)

func (s *Server) SetupReadHandler() {
	s.Websocket.HandleMessage(func(session *melody.Session, msg []byte) {
		log.Println("Received message from", session.Request.RemoteAddr, ":", string(msg))

		if string(msg) == "world" && world.IsWorldWaiting {
			log.Println("World was received")
			world.IsWorldWaiting = false
		}
	})
}

func (s *Server) SendToClients(sender *melody.Session, msg []byte) error {
	return s.Websocket.BroadcastFilter(msg, func(session *melody.Session) bool {
		return sender != session
	})
}
