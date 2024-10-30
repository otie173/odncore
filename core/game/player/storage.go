package player

import (
	"log"
	"os"

	"github.com/otie173/odncore/utils/config"
	"github.com/otie173/odncore/utils/filesystem"
	"github.com/otie173/odncore/utils/logger"
	"github.com/vmihailenco/msgpack/v5"
)

func InitPlayer(cfg config.Config) {
	dirs := []string{filesystem.PLAYERS_DIR_PATH, filesystem.PLAYER_DATA_DIR_PATH, filesystem.PLAYER_DB_PATH}

	for _, path := range dirs {
		if !filesystem.DirExists(path) {
			err := os.Mkdir(path, 0755)
			if err != nil {
				logger.Error("Error creating directory: ", err)
			}
		}
	}

	players = make(map[string]Player, cfg.MaxPlayers)
}

func GetName(addr string) string {
	log.Println(addr)
	log.Println(players)

	if player, ok := players[addr]; ok {
		log.Println("Player found:", player)
		log.Println("Nickname:", player.nickname)
	} else {
		log.Println("Player not found for address:", addr)
	}

	return players[addr].nickname
}

func AddPlayer(addr string, nickname string) {
	players[addr] = Player{nickname: nickname, inventory: Inventory{}}
}

func Save(nickname string) error {
	data, err := msgpack.Marshal(players)
	if err != nil {
		return err
	}

	os.WriteFile(filesystem.PLAYER_DATA_DIR_PATH+"inventory.odn", data, 0644)
	logger.Info("Inventory saved succesfully")
	return nil
}

func Load(nickname string) {

}
