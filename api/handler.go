package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/otie173/odncore/internal/auth"
	"github.com/otie173/odncore/internal/game/player"
	"github.com/otie173/odncore/internal/game/world"
	"github.com/otie173/odncore/internal/server"
	"github.com/otie173/odncore/internal/utils/config"
	"github.com/otie173/odncore/internal/utils/database"
	"github.com/otie173/odncore/internal/utils/filesystem"
	"github.com/otie173/odncore/internal/utils/logger"
)

// @Summary Get server status
// @Description Get current server status
// @Tags server
// @Accept json
// @Produce json
// @Success 200 {object} server.ServerStatus
// @Router /server/status [get]
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := server.GetStatus()
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// @Summary Get server info
// @Description Get information about the server
// @Tags server
// @Accept json
// @Produce json
// @Success 200 {object} server.ServerInfo
// @Router /server/info [get]
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	info := server.GetInfo(config.GetConfig())
	if err := json.NewEncoder(w).Encode(info); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// @Summary Authenticate player
// @Description Register new player or login existing player
// @Tags player
// @Accept json
// @Produce plain
// @Param playerAuth body auth.PlayerAuth true "Player authentication data"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "FAIL"
// @Router /player/auth [post]
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
	}
}

// @Summary Get player data
// @Description Get player's saved data
// @Tags player
// @Accept json
// @Produce json
// @Param Session-Nickname header string true "Player nickname"
// @Success 200 {object} player.Player
// @Failure 404 {string} string "Player not found"
// @Router /player/getpdata [get]
func GetPDataHandler(w http.ResponseWriter, r *http.Request) {
	playerNickname := r.Header.Get("Session-Nickname")

	if filesystem.FileExists(filesystem.PLAYER_DATA_DIR_PATH + playerNickname + ".odn") {
		playerData, err := player.Load(playerNickname)
		if err != nil {
			logger.Error("Error with load player data: ", err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(playerData)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

// @Summary Save player data
// @Description Save player's current state
// @Tags player
// @Accept json
// @Param Session-Nickname header string true "Player nickname"
// @Param playerData body player.Player true "Player data to save"
// @Success 200
// @Router /player/loadpdata [post]
func LoadPDataHandler(w http.ResponseWriter, r *http.Request) {
	playerNickname := r.Header.Get("Session-Nickname")

	playerData, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("Error with read request body: %v", err)
	}
	defer r.Body.Close()

	if err := player.Save(playerNickname, playerData); err != nil {
		logger.Error("Error with save player data: ", err)
	}
}

// @Summary Load world ID
// @Description Load world identification data
// @Tags world
// @Accept json
// @Param idData body string true "World ID data"
// @Success 200
// @Router /world/loadid [post]
func LoadIdHandler(w http.ResponseWriter, r *http.Request) {
	idData, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("Error with read request body: %v", err)
	}
	defer r.Body.Close()

	if err := world.LoadIdNetwork(idData); err != nil {
		logger.Error("Error with load id from network: ", err)
	}
}

// @Summary Get world data
// @Description Get current world state as binary data
// @Tags world
// @Produce octet-stream
// @Success 200 {string} binary "World binary data"
// @Router /world/getworld [get]
func GetWorldHandler(w http.ResponseWriter, r *http.Request) {
	world.Save()

	worldData, err := os.ReadFile(filesystem.WORLD_DIR_PATH + "world.odn")
	if err != nil {
		logger.Errorf("Error with read world file: %v", err)
	}

	w.Write(worldData)
}

// @Summary Load world data
// @Description Load world state data from binary file
// @Tags world
// @Accept octet-stream
// @Param worldData body string true "World binary data" format(binary)
// @Success 200
// @Router /world/loadworld [post]
func LoadWorldHandler(w http.ResponseWriter, r *http.Request) {
	worldData, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("Error with read request body: %v", err)
	}
	defer r.Body.Close()

	if err := world.ByteToFile(worldData); err != nil {
		logger.Error("Error with convert world bytes to file: ", err)
	}
}
