package api

import (
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("GET /api/info", InfoHandler)
	http.HandleFunc("GET /api/status", StatusHandler)
	http.HandleFunc("GET /api/getpdata", GetPDataHandler)
	http.HandleFunc("GET /api/getworld", GetWorldHandler)

	http.HandleFunc("POST /api/auth", AuthHandler)
	http.HandleFunc("POST /api/loadid", LoadIdHandler)
	http.HandleFunc("POST /api/loadworld", LoadWorldHandler)
}
