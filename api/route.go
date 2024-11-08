package api

import (
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("GET /api/status", StatusHandler)
	http.HandleFunc("GET /api/getworld", GetWorldHandler)
	http.HandleFunc("GET /api/loadworld", LoadWorldHandler)
	http.HandleFunc("POST /api/auth", AuthHandler)
}
