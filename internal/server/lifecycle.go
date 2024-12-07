package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/olahol/melody"
	"github.com/otie173/odncore/internal/game/player"
	"github.com/otie173/odncore/internal/utils/logger"
)

func Start(r *chi.Mux) {
	r.HandleFunc("GET /ws", func(res http.ResponseWriter, req *http.Request) {
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
		logger.Player(sessionNickname, "joined the game")
	})

	websocket.HandleDisconnect(func(session *melody.Session) {
		sessionNickname := session.Request.Header.Get("Session-Nickname")
		rejected, _ := session.Get("rejected")
		if rejected == nil && playersConnected > 0 {
			playersConnected--
			player.Remove(sessionNickname)
			if !session.IsClosed() {
				if err := sendPlayersList(); err != nil {
					logger.Error("Failed to send players list to clients: ", err)
				}
			}
			logger.Player(sessionNickname, "left the game")
		}
	})
	logger.Info("Server started on address", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		logger.Fatal("Failed to start server: ", err)
	}
}

func Stop() error {
	if err := websocket.Close(); err != nil {
		return err
	}
	return nil
}
