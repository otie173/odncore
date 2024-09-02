package player

import (
	"log"
	"os"
)

func InitPlayer() {
	if !dirExists() {
		err := os.Mkdir("./players", 0755)
		if err != nil {
			log.Println("Error creating directory: ", err)
		} else {
			log.Println("Players directory created successfully")
		}
	}
}

func dirExists() bool {
	_, err := os.Stat("./players")
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
