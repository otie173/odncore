package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/otie173/odncore/api"
	"github.com/otie173/odncore/internal/game/player"
	"github.com/otie173/odncore/internal/game/world"
	"github.com/otie173/odncore/internal/server"
	"github.com/otie173/odncore/internal/utils/config"
	"github.com/otie173/odncore/internal/utils/database"
	"github.com/otie173/odncore/internal/utils/filesystem"
	"github.com/otie173/odncore/internal/utils/logger"
	"github.com/otie173/odncore/internal/utils/webhook/discord"
)

func initSystem() config.Config {
	logger.InitLogger()

	config.InitConfig()
	if err := config.Load(); err != nil {
		logger.Fatal("Failed to load config: ", err)
	}
	cfg := config.GetConfig()

	if err := database.InitDB(); err != nil {
		logger.Fatal("Failed to create database: ", err)
	}
	return cfg
}

func initIntegrations() {
	discord.InitDiscord()
}

func initGame(cfg config.Config) {
	if err := world.InitWorld(); err != nil {
		logger.Fatal("Failed to init world: ", err)
	}

	if err := player.InitPlayer(cfg); err != nil {
		logger.Fatal("Failed init player: ", err)
	}
}

func initCore(cfg config.Config) {
	server.InitServer(cfg.Address, cfg.MaxPlayers)
	server.InitHandler()
	if filesystem.FileExists(filesystem.WORLD_DIR_PATH + "id.odn") {
		if err := world.LoadIdFile(); err != nil {
			logger.Fatal("Failed to load id file for world: ", err)
		}
	}
	if filesystem.FileExists(filesystem.WORLD_DIR_PATH + "world.odn") {
		if err := world.Load(); err != nil {
			logger.Fatal("Failed to load world: ", err)
		}
	}
}

func initServer() {
	go func() {
		r := api.InitRoutes()
		server.Start(r)
	}()
	logger.Server("Server is running")
}

func shutdownServer() {
	logger.Server("Shutting down server")
	if err := server.Stop(); err != nil {
		logger.Error("Failed to stop the server: ", err)
	}

	if err := config.Save(); err != nil {
		logger.Error("Failed to save config: ", err)
	}

	if err := database.Save(); err != nil {
		logger.Error("Failed save database: ", err)
	}
	if !world.IsIdWaiting {
		if err := world.SaveId(); err != nil {
			logger.Error("Failed to save id file: ", err)
		}
	}
	if !world.IsWorldWaiting {
		if err := world.Save(); err != nil {
			logger.Error("Failed to save world: ", err)
		}
	}
}

func main() {
	cfg := initSystem()
	initIntegrations()
	initGame(cfg)
	initCore(cfg)
	initServer()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	shutdownServer()
}
