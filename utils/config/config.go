package config

import (
	"encoding/json"
	"log"
	"os"
)

var (
	cfg *Config
)

const CONFIG_PATH string = "config.json"

type Config struct {
	Address    string `json:"address"`
	MaxPlayers int    `json:"max_player"`
}

func NewConfig() *Config {
	return &Config{
		Address:    "0.0.0.0:8080",
		MaxPlayers: 16,
	}
}

func (c *Config) Load() error {
	if _, err := os.Stat(CONFIG_PATH); os.IsNotExist(err) {
		return c.Save()
	}

	file, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	return json.Unmarshal(file, c)
}

func (c *Config) Save() error {
	data, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		log.Println("Error saving config: ", err)
	} else {
		log.Println("Config saved successfully")
	}

	return os.WriteFile(CONFIG_PATH, data, 0644)
}
