package config

import (
	"encoding/json"
	"os"
)

var (
	Cfg Config
)

const CONFIG_PATH string = "config.json"

type Config struct {
	Address               string `json:"address"`
	MaxPlayers            int    `json:"max_player"`
	DiscordWebhookEnabled bool   `json:"discord_webhook_enabled"`
	DiscordWebhookURL     string `json:"discord_webhook_url"`
	DiscordWebhookName    string `json:"discord_webhook_name"`
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
		return err
	}
	return json.Unmarshal(file, &Cfg)
}

func Save() error {
	data, err := json.MarshalIndent(Cfg, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(CONFIG_PATH, data, 0644)
}
