package database

import (
	"github.com/otie173/odncore/utils/logger"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	db *leveldb.DB
)

const (
	PLAYER_DATA_PATH string = "./players/db/"
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
		logger.Error("Error saving database: ", err)
	} else {
		logger.Error("Database saved successfully")
	}
}

func AddPlayer(nickname string, password string) {
	db.Put([]byte(nickname), []byte(password), nil)
}

func GetPasswordHash(nickname string) ([]byte, error) {
	password, err := db.Get([]byte(nickname), nil)
	if err != nil {
		return nil, err
	}
	return password, err
}

func PlayerExists(nickname string) bool {
	if exists, _ := db.Has([]byte(nickname), nil); exists {
		return true
	}
	return false
}
