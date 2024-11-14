package player

import (
	"log"
	"os"

	"github.com/otie173/odncore/internal/utils/config"
	"github.com/otie173/odncore/internal/utils/filesystem"
	"github.com/vmihailenco/msgpack/v5"
)

func InitPlayer(cfg config.Config) error {
	dirs := []string{filesystem.PLAYERS_DIR_PATH, filesystem.PLAYER_DATA_DIR_PATH, filesystem.PLAYER_DB_PATH}

	for _, path := range dirs {
		if !filesystem.DirExists(path) {
			err := os.Mkdir(path, 0755)
			if err != nil {
				return err
			}
		}
	}

	players = make(map[string]Player, cfg.MaxPlayers)
	return nil
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

func Save(addr string) error {
	data, err := msgpack.Marshal(players[addr])
	if err != nil {
		return err
	}

	os.WriteFile(filesystem.PLAYER_DATA_DIR_PATH+players[addr].nickname+".odn", data, 0644)
	return nil
}

func Load(nickname string) error {
	return nil
}
