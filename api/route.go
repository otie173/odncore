package api

import (
	"net/http"
)

func SetupRoutes() {
	// Handlers about server
	http.HandleFunc("GET /api/info", InfoHandler)
	http.HandleFunc("GET /api/status", StatusHandler)

	// Handlers about world
	http.HandleFunc("GET /api/getworld", GetWorldHandler)
	http.HandleFunc("POST /api/loadid", LoadIdHandler)
	http.HandleFunc("POST /api/loadworld", LoadWorldHandler)

	// Handlers about player
	http.HandleFunc("GET /api/getplayerdata", GetPDataHandler)
	http.HandleFunc("POST /api/auth", AuthHandler)
	http.HandleFunc("POST /api/loadplayerdata", LoadPDataHandler)
}
