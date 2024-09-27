package database

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	db *leveldb.DB
)

const (
	PLAYER_DATA_PATH string = "./players/db"
)

func NewDatabase() error {
	var err error
	db, err = leveldb.OpenFile(PLAYER_DATA_PATH, nil)
	if err != nil {
		return err
	}
	return nil
}

func Close() error {
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}

func Save() {
	if err := Close(); err != nil {
		log.Println("Error saving database: ", err)
	} else {
		log.Println("Database saved successfully")
	}
}

func AddPlayer(nickname string, password string) {
	db.Put([]byte(nickname), []byte(password), nil)
}

func PlayerExists(nickname string) bool {
	if exists, _ := db.Has([]byte(nickname), nil); exists {
		return true
	}
	return false
}
