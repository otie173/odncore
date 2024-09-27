package api

import (
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("GET /api/about", AboutHandler())
	http.HandleFunc("POST /api/auth", AuthHandler())
}
