package setup

import (
	"github.com/otie173/odncore/core/network/types"
)

type API interface {
	GetServer() types.ServerInterface
}

type api struct {
	server types.ServerInterface
}

func NewAPI(server types.ServerInterface) API {
	return &api{server: server}
}

func (a *api) GetServer() types.ServerInterface {
	return a.server
}
