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

	s.Websocket.HandleMessageBinary(func(s *melody.Session, b []byte) {
		if world.IsWorldWaiting {
			log.Println("World was received")
			world.IsWorldWaiting = false

			if err := world.ByteToFile(b); err != nil {
				log.Println("Error: ", err)
			}
		}
	})
}

func (s *Server) SendRequest(session *melody.Session, opcode byte) error {
	if err := session.WriteBinary([]byte{opcode}); err != nil {
		return err
	}

	return nil
}

func (s *Server) SendToClients(sender *melody.Session, msg []byte) error {
	return s.Websocket.BroadcastFilter(msg, func(session *melody.Session) bool {
		return sender != session
	})
}
