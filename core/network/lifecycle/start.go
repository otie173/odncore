package lifecycle

import (
	"net/http"

	"github.com/olahol/melody"
	"github.com/otie173/odncore/core/network/types"
	"github.com/otie173/odncore/utils/logger"
)

func Start(s *types.Server) error {
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
	})

	s.Websocket.HandleDisconnect(func(session *melody.Session) {
		rejected, _ := session.Get("rejected")
		if rejected == nil && s.PlayersConnected > 0 {
			s.PlayersConnected--
			logger.PlayerDisconnected(session.Request.RemoteAddr)
		}
	})

	logger.StartServer(s.Addr)
	return http.ListenAndServe(s.Addr, nil)
}
