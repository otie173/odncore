package route

import (
	"net/http"

	"github.com/otie173/odncore/api/handler"
	"github.com/otie173/odncore/core/server"
)

func SetupRoutes(s *server.Server) {
	http.HandleFunc("GET /api/about", handler.AboutHandler(s))
}
