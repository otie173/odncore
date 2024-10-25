package player

import (
	"os"

	"github.com/otie173/odncore/utils/config"
	"github.com/otie173/odncore/utils/filesystem"
	"github.com/otie173/odncore/utils/logger"
	"github.com/vmihailenco/msgpack/v5"
)

func InitPlayer(cfg config.Config) {
	dirs := []string{filesystem.PLAYERS_DIR_PATH, filesystem.PLAYER_DATA_DIR_PATH, filesystem.PLAYER_DB_PATH}

	for _, path := range dirs {
		if !dirExists(path) {
			err := os.Mkdir(path, 0755)
			if err != nil {
				logger.Error("Error creating directory: ", err)
			}
		}
	}

	players = make([]Player, cfg.MaxPlayers)
}

func dirExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func AddPlayer(nickname string) {
	players = append(players, Player{nickname: nickname, inventory: Inventory{}})
}

func InventorySave() error {
	data, err := msgpack.Marshal(players)
	if err != nil {
		return err
	}

	os.WriteFile(filesystem.PLAYER_DATA_DIR_PATH+"inventory.odn", data, 0644)
	logger.Info("Inventory saved succesfully")
	return nil
}

func InventoryLoad() {

}
