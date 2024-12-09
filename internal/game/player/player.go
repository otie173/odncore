package player

import (
	"net"
	"os"

	"github.com/otie173/odncore/internal/utils/config"
	"github.com/otie173/odncore/internal/utils/filesystem"
)

var (
	players map[net.Addr]string
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

	players = make(map[net.Addr]string, cfg.MaxPlayers)
	return nil
}

func Add(addr net.Addr, nickname string) {
	players[addr] = nickname
}

func Remove(addr net.Addr) {
	delete(players, addr)
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
