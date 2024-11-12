package server

import (
	"net/http"

	"github.com/olahol/melody"
	"github.com/otie173/odncore/internal/utils/logger"
	"github.com/otie173/odncore/internal/utils/webhook/discord"
)

func Start() {
	http.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request) {
		websocket.HandleRequest(res, req)
	})

	websocket.HandleConnect(func(session *melody.Session) {
		sessionNickname := session.Request.Header.Get("Session-Nickname")

		if playersConnected >= maxPlayers {
			session.Write([]byte("Sorry! Server is full"))
			session.Set("rejected", true)
			session.Close()
			return
		}
		playersConnected++
		logger.Info("Player connected: " + sessionNickname)

		if discord.WebhookEnabled() {
			discord.PlayerMessage("Player connected:", sessionNickname)
		}
	})

	websocket.HandleDisconnect(func(session *melody.Session) {
		sessionNickname := session.Request.Header.Get("Session-Nickname")

		rejected, _ := session.Get("rejected")
		if rejected == nil && playersConnected > 0 {
			playersConnected--
			logger.Info("Player disconnected: " + sessionNickname)

			if discord.WebhookEnabled() {
				discord.PlayerMessage("Player disconnected:", sessionNickname)
			}
		}
	})

	logger.Info("Server started on address", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logger.Fatal("Failed to start server: ", err)
	}
}

func Stop() {
	websocket.Close()
	logger.Info("Server was stopped")
}
