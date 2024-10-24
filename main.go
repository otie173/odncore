package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/otie173/odncore/api"
	"github.com/otie173/odncore/core/game/player"
	"github.com/otie173/odncore/core/game/world"
	"github.com/otie173/odncore/core/server"
	"github.com/otie173/odncore/utils/config"
	"github.com/otie173/odncore/utils/database"
	"github.com/otie173/odncore/utils/logger"
)

func init() {
	// Init server things
	logger.Register()

	config.NewConfig()
	config.Load()

	if err := database.NewDatabase(); err != nil {
		logger.Fatal("Error: ", err)
	}

	// Init game things
	world.InitWorld()
	player.InitPlayer(config.Cfg)
}

func run() {
	server.New(config.Cfg.Address, config.Cfg.MaxPlayers)
	server.SetupReadHandler()
	api.SetupRoutes()
	if world.FileExists(world.WORLD_DIR_PATH + "id.odn") {
		world.LoadIdFile()
	}
	if world.FileExists(world.WORLD_DIR_PATH + "world.odn") {
		world.Load()
	}

	go func() {
		server.Start()
	}()
	logger.Info("Server is running. Press CTRL+C to stop.")

	// Graceful shutdown
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
	player.InventorySave()
}

func main() {
	run()
}
