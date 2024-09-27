package server

import (
	"log"
	"net/http"

	"github.com/olahol/melody"
	"github.com/otie173/odncore/core/game/world"
	"github.com/otie173/odncore/utils/logger"
)

func Start() {
	http.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request) {
		websocket.HandleRequest(res, req)
	})

	websocket.HandleConnect(func(session *melody.Session) {
		if playersConnected >= maxPlayers {
			//session.Write([]byte("Sorry! Server is full"))
			session.Set("rejected", true)
			session.Close()
			return
		}
		playersConnected++
		logger.PlayerConnected(session.Request.RemoteAddr)

		if world.IsWorldWaiting {
			if playersConnected == 1 {
				if err := ReceiveWorld(session); err != nil {
					log.Fatal("Fail with receive world from client: ", err)
				}
			}
		} else {
			if err := SendWorld(session); err != nil {
				log.Fatal("Fail with send world to client: ", err)
			}
		}
	})

	websocket.HandleDisconnect(func(session *melody.Session) {
		rejected, _ := session.Get("rejected")
		if rejected == nil && playersConnected > 0 {
			playersConnected--
			logger.PlayerDisconnected(session.Request.RemoteAddr)
		}
	})

	logger.StartServer(addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func Stop() {
	websocket.Close()
	logger.StopServer()
}
