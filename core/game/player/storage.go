package player

import (
	"log"
	"os"
)

const (
	PLAYERS_DIR_PATH     string = "./players"
	PLAYER_DATA_DIR_PATH string = "./players/data"
	PLAYER_DB_PATH       string = "./players/db"
)

func InitPlayer() {
	dirs := []string{PLAYERS_DIR_PATH, PLAYER_DATA_DIR_PATH, PLAYER_DB_PATH}

	for _, path := range dirs {
		if !dirExists(path) {
			err := os.Mkdir(path, 0755)
			if err != nil {
				log.Println("Error creating directory: ", err)
			}
		}
	}
}

func dirExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
