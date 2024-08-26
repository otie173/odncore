package api

import (
	"net/http"

	"github.com/otie173/odncore/core/server"
)

func SetupRoutes(s *server.Server) {
	http.HandleFunc("GET /api/about", AboutHandler(s))
}
