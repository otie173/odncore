package server

import (
	"log"
	"net/http"

	"github.com/olahol/melody"
	"github.com/otie173/odncore/core/game/world"
	"github.com/otie173/odncore/utils/logger"
)

func (s *Server) Start() {
	http.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request) {
		s.Websocket.HandleRequest(res, req)
	})

	s.Websocket.HandleConnect(func(session *melody.Session) {
		if s.PlayersConnected >= s.MaxPlayers {
			session.Write([]byte("Sorry! Server is full"))
			session.Set("rejected", true)
			session.Close()
			return
		}
		s.PlayersConnected++
		logger.PlayerConnected(session.Request.RemoteAddr)

		if world.IsWorldWaiting {
			if s.PlayersConnected == 1 {
				log.Println("Wait a world from client side")
				if err := s.ReceiveWorld(session); err != nil {
					log.Println("Fail with receive world from client: ", err)
				}
			}
		} else {
			if err := s.SendWorld(session); err != nil {
				log.Println("Fail with send world to client: ", err)
			}
		}
	})

	s.Websocket.HandleDisconnect(func(session *melody.Session) {
		rejected, _ := session.Get("rejected")
		if rejected == nil && s.PlayersConnected > 0 {
			s.PlayersConnected--
			logger.PlayerDisconnected(session.Request.RemoteAddr)
		}
	})

	logger.StartServer(s.Addr)
	if err := http.ListenAndServe(s.Addr, nil); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func (s *Server) Stop() {
	s.Websocket.Close()
	logger.StopServer()
}
