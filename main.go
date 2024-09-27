package main

import (
	"log"
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
		log.Fatalf("Error: %v", err)
	}

	// Init game things
	world.InitWorld()
	player.InitPlayer()
}

func run() {
	server.New(config.Cfg.Address, config.Cfg.MaxPlayers)
	server.SetupReadHandler()
	api.SetupRoutes()
	if world.WorldExists() {
		world.Load()
	}

	go func() {
		server.Start()
	}()
	log.Println("Server is running. Press CTRL+C to stop.")

	// Graceful shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutting down server")
	server.Stop()

	config.Save()
	database.Save()
	if !world.IsWorldWaiting {
		world.Save()
	}
}

func main() {
	run()
}
