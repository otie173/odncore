package player

import (
	"os"

	"github.com/otie173/odncore/internal/utils/config"
	"github.com/otie173/odncore/internal/utils/filesystem"
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

func Add(nickname string, x, y, targetX, targetY float32) {
	players[nickname] = Player{nickname, x, y, targetX, targetY}
}

func Remove(nickname string) {
	delete(players, nickname)
}

func GetList() map[string]Player {
	return players
}

func Save(nickname string, data []byte) error {
	if err := os.WriteFile(filesystem.PLAYER_DATA_DIR_PATH+nickname+".odn", data, 0644); err != nil {
		return err
	}
	return nil
}

func Load(nickname string) ([]byte, error) {
	data, err := os.ReadFile(filesystem.PLAYER_DATA_DIR_PATH + nickname + ".odn")
	if err != nil {
		return nil, err
	}
	return data, nil
}
