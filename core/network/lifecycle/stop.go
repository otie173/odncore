package lifecycle

import (
	"github.com/otie173/odncore/core/network/server"
	"github.com/otie173/odncore/utils/logger"
)

func Stop(s *server.Server) error {
	s.Websocket.Close()
	logger.StopServer()
	return nil
}
