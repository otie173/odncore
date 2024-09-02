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
	"github.com/otie173/odncore/utils/logger"
)

func init() {
	world.InitWorld()
	player.InitPlayer()
}

func run(cfg config.Config) {
	server := server.New(cfg.Address, cfg.MaxPlayers)
	server.SetupReadHandler()
	api.SetupRoutes(server)
	if world.WorldExists() {
		world.Load()
	}

	go func() {
		server.Start()
	}()
	log.Println("Server is running. Press CTRL+C to stop.")

	// Graceful shutdownc
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutting down server")
	server.Stop()

	cfg.Save()
	world.Save()
}

func main() {
	logger.Register()

	cfg := config.NewConfig()
	cfg.Load()

	run(*cfg)
}
