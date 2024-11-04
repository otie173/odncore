package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/otie173/odncore/api"
	"github.com/otie173/odncore/internal/config"
	"github.com/otie173/odncore/internal/database"
	"github.com/otie173/odncore/internal/filesystem"
	"github.com/otie173/odncore/internal/game/player"
	"github.com/otie173/odncore/internal/game/world"
	"github.com/otie173/odncore/internal/logger"
	"github.com/otie173/odncore/internal/server"
)

func init() {
	logger.Register()

	config.NewConfig()
	config.Load()

	if err := database.NewDatabase(); err != nil {
		logger.Fatal("Error: ", err)
	}

	world.InitWorld()
	player.InitPlayer(config.Cfg)
}

func run() {
	server.New(config.Cfg.Address, config.Cfg.MaxPlayers)
	server.SetupReadHandler()
	api.SetupRoutes()
	if filesystem.FileExists(filesystem.WORLD_DIR_PATH + "id.odn") {
		world.LoadIdFile()
	}
	if filesystem.FileExists(filesystem.WORLD_DIR_PATH + "world.odn") {
		world.Load()
	}

	go func() {
		server.Start()
	}()
	logger.Info("Server is running. Press CTRL+C to stop.")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	logger.Info("Shutting down server")
	server.Stop()

	config.Save()
	database.Save()
	if !world.IsIdWaiting {
		world.SaveId()
	}
	if !world.IsWorldWaiting {
		world.Save()
	}
}

func main() {
	run()
}
