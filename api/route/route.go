package route

import (
	"net/http"

	"github.com/otie173/odncore/api/handler"
	"github.com/otie173/odncore/api/setup"
)

func SetupRoutes(api setup.API) {
	http.HandleFunc("GET /api/about", handler.AboutHandler(api))
}
