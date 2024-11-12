package api

import (
	"net/http"
)

func SetupRoutes() {
	// GET Methods
	http.HandleFunc("GET /api/status", StatusHandler)
	http.HandleFunc("GET /api/getworld", GetWorldHandler)

	// POST Methods
	http.HandleFunc("POST /api/auth", AuthHandler)
	http.HandleFunc("POST /api/loadid", LoadIdHandler)
	http.HandleFunc("POST /api/loadworld", LoadWorldHandler)
}
