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

func main() {
	logger.Register()

	config.NewConfig()
	if err := config.Load(); err != nil {
		logger.Fatal("Error with load config: ", err)
	}
	cfg := config.GetConfig()

	if err := database.NewDatabase(); err != nil {
		logger.Fatal("Error with create database: ", err)
	}

	discord.InitDiscord()

	if err := world.InitWorld(); err != nil {
		logger.Fatal("Error with init world: ", err)
	}

	if err := player.InitPlayer(cfg); err != nil {
		logger.Fatal("Error with init player: ", err)
	}

	server.New(cfg.Address, cfg.MaxPlayers)
	server.SetupReadHandler()
	api.SetupRoutes()
	if filesystem.FileExists(filesystem.WORLD_DIR_PATH + "id.odn") {
		if err := world.LoadIdFile(); err != nil {
			logger.Fatal("Error with load id file for world: ", err)
		}
	}
	if filesystem.FileExists(filesystem.WORLD_DIR_PATH + "world.odn") {
		if err := world.Load(); err != nil {
			logger.Fatal("Error with load world: ", err)
		}
	}

	go func() {
		server.Start()
	}()
	logger.Server("Server is running")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	logger.Server("Shutting down server")
	if err := server.Stop(); err != nil {
		logger.Error("Error with stop the server: ", err)
	}

	if err := config.Save(); err != nil {
		logger.Error("Error with save config: ", err)
	}

	if err := database.Save(); err != nil {
		logger.Error("Error with save database: ", err)
	}
	if !world.IsIdWaiting {
		if err := world.SaveId(); err != nil {
			logger.Error("Error with save id file: ", err)
		}
	}
	if !world.IsWorldWaiting {
		if err := world.Save(); err != nil {
			logger.Error("Errorw with save world: ", err)
		}
	}
}
