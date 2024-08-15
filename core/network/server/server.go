package server

import (
	"github.com/olahol/melody"
	"github.com/otie173/odncore/core/network/info"
	"github.com/otie173/odncore/core/network/io"
	"github.com/otie173/odncore/core/network/lifecycle"
	"github.com/otie173/odncore/core/network/types"
)

type Server struct {
	types.Server
}

func New(addr string, maxPlayers int) types.ServerInterface {
	return &Server{
		Server: types.Server{
			Websocket:  melody.New(),
			Addr:       addr,
			MaxPlayers: maxPlayers,
		},
	}
}

func (s *Server) Start() error {
	return lifecycle.Start(&s.Server)
}

func (s *Server) Stop() error {
	return lifecycle.Stop(&s.Server)
}

func (s *Server) GetInfo() types.ServerInfo {
	return info.GetInfo(&s.Server)
}

func (s *Server) SetupReadHandler() {
	io.SetupReadHandler(&s.Server, s.SendToClients)
}

func (s *Server) SendToClients(sender *melody.Session, msg []byte) error {
	return io.SendToClients(&s.Server, sender, msg)
}
