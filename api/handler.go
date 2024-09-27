package api

import (
	"encoding/json"
	"net/http"

	"github.com/otie173/odncore/core/auth"
	"github.com/otie173/odncore/core/server"
	"github.com/otie173/odncore/utils/database"
)

func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AboutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info := server.GetInfo()
		respondJSON(w, info)
	}
}

func AuthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var player auth.Player

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&player); err != nil {
			http.Error(w, "Error parsing JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		switch database.PlayerExists(player.Nickname) {
		case false:
			auth.RegisterPlayer(player.Nickname, player.Password)
		case true:
			auth.LoginPlayer(player.Nickname, player.Password)
		}
	}
}
