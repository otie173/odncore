package io

import (
	"log"

	"github.com/olahol/melody"
	"github.com/otie173/odncore/core/network/types"
)

func SetupReadHandler(s *types.Server, sendFunc func(*melody.Session, []byte) error) {
	s.Websocket.HandleMessage(func(session *melody.Session, msg []byte) {
		err := sendFunc(session, msg)
		if err != nil {
			log.Println("Error sending message: ", err)
		}
		log.Println("Received message from", session.Request.RemoteAddr, ":", string(msg))
	})
}
