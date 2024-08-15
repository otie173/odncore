package lifecycle

import (
	"github.com/otie173/odncore/core/network/types"
	"github.com/otie173/odncore/utils/logger"
)

func Stop(s *types.Server) error {
	s.Websocket.Close()
	logger.StopServer()
	return nil
}
