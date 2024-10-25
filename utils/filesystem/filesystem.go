package filesystem

import "os"

const (
	PLAYERS_DIR_PATH     string = "./players/"
	PLAYER_DATA_DIR_PATH string = "./players/data/"
	PLAYER_DB_PATH       string = "./players/db/"

	WORLD_DIR_PATH string = "./world/"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func DirExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
