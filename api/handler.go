package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/otie173/odncore/internal/auth"
	"github.com/otie173/odncore/internal/game/player"
	"github.com/otie173/odncore/internal/game/world"
	"github.com/otie173/odncore/internal/server"
	"github.com/otie173/odncore/internal/utils/database"
	"github.com/otie173/odncore/internal/utils/filesystem"
	"github.com/otie173/odncore/internal/utils/logger"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := server.GetStatus()
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetWorldHandler(w http.ResponseWriter, r *http.Request) {
	world.Save()

	worldData, err := os.ReadFile(filesystem.WORLD_DIR_PATH + "world.odn")
	if err != nil {
		logger.Errorf("Error with read world file: %v", err)
	}

	w.Write(worldData)
}

func LoadWorldHandler(w http.ResponseWriter, r *http.Request) {
	bodyData, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("Error with read request body: %v", err)
	}
	defer r.Body.Close()

	log.Println(bodyData)
	w.Write(bodyData)
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
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
