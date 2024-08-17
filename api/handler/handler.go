package handler

import (
	"encoding/json"
	"net/http"

	"github.com/otie173/odncore/core/server"
)

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AboutHandler(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info := s.GetInfo()
		respondJSON(w, info)
	}
}
