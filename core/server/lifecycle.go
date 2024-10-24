package server

import (
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
			session.Write([]byte("Sorry! Server is full"))
			session.Set("rejected", true)
			session.Close()
			return
		}
		playersConnected++
		logger.Info("Player connected: ", session.Request.RemoteAddr)

		if world.IsIdWaiting {
			if playersConnected == 1 {
				if err := AskId(session); err != nil {
					logger.Fatal("Fail with ask id from client: ", err)
				}
			}
		}
		if world.IsWorldWaiting {
			if playersConnected == 1 {
				if err := AskWorld(session); err != nil {
					logger.Fatal("Fail with ask world from client: ", err)
				}
			}
		} else {
			if err := SendWorld(session); err != nil {
				logger.Fatal("Fail with send world to client: ", err)
			}
		}
	})

	websocket.HandleDisconnect(func(session *melody.Session) {
		rejected, _ := session.Get("rejected")
		if rejected == nil && playersConnected > 0 {
			playersConnected--
			logger.Info("Player disconnected: ", session.Request.RemoteAddr)
		}
	})

	logger.Info("Server started on ", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logger.Fatal("Failed to start server: ", err)
	}
}

func Stop() {
	websocket.Close()
	logger.Info("Server was stopped")
}
