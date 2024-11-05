package api

import (
	"encoding/json"
	"net/http"

	"github.com/otie173/odncore/internal/auth"
	"github.com/otie173/odncore/internal/game/player"
	"github.com/otie173/odncore/internal/server"
	"github.com/otie173/odncore/internal/utils/database"
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
		var playerAuth auth.PlayerAuth

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&playerAuth); err != nil {
			http.Error(w, "Error parsing JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		var authorizationOK bool
		switch database.PlayerExists(playerAuth.Nickname) {
		case false:
			authorizationOK = auth.RegisterPlayer(playerAuth.Nickname, playerAuth.Password)
		case true:
			authorizationOK = auth.LoginPlayer(playerAuth.Nickname, playerAuth.Password)
		}

		switch authorizationOK {
		case false:
			w.Write([]byte("FAIL"))
		case true:
			w.Write([]byte("OK"))

			player.AddPlayer(r.RemoteAddr, playerAuth.Nickname)
		}
	}
}
