package handler

import (
	"encoding/json"
	"net/http"

	"github.com/otie173/odncore/core/network/info"
	"github.com/otie173/odncore/core/network/server"
)

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AboutHandler(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info := info.GetInfo(s)
		respondJSON(w, info)
	}
}
