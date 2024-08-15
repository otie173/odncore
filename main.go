package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/otie173/odncore/api/route"
	"github.com/otie173/odncore/api/setup"
	"github.com/otie173/odncore/core/network/server"
	"github.com/otie173/odncore/utils/config"
	"github.com/otie173/odncore/utils/logger"
)

func main() {
	logger.Register()

	cfg := config.NewConfig()
	err := cfg.Load()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	server := server.New(cfg.Address, cfg.MaxPlayers)
	server.SetupReadHandler()

	api := setup.NewAPI(server)
	route.SetupRoutes(api)

	go func() {
		if err := server.Start(); err != nil {
			log.Fatal("Failed to start server: ", err)
		}
	}()

	log.Println("Server is running. Press CTRL+C to stop.")

	// Graceful shutdownc
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutting down server")
	if err := server.Stop(); err != nil {
		log.Println("Error during server shutdown:", err)
	}

	if err := cfg.Save(); err != nil {
		log.Println("Error saving config: ", err)
	} else {
		log.Println("Config saved successfully")
	}
}
