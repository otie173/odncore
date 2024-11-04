package config

import (
	"encoding/json"
	"os"

	"github.com/otie173/odncore/internal/logger"
)

var (
	Cfg Config
)

const CONFIG_PATH string = "config.json"

type Config struct {
	Address    string `json:"address"`
	MaxPlayers int    `json:"max_player"`
}

func NewConfig() {
	Cfg = Config{
		Address:    "0.0.0.0:8080",
		MaxPlayers: 16,
	}
}

func Load() error {
	if _, err := os.Stat(CONFIG_PATH); os.IsNotExist(err) {
		return Save()
	}

	file, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		logger.Fatal("Failed to load config: ", err)
	}

	return json.Unmarshal(file, &Cfg)
}

func Save() error {
	data, err := json.MarshalIndent(Cfg, "", " ")
	if err != nil {
		logger.Error("Error saving config: ", err)
	} else {
		logger.Info("Config saved successfully")
	}

	return os.WriteFile(CONFIG_PATH, data, 0644)
}
