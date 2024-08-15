package logger

import (
	"log"
	"os"
)

var (
	logger *log.Logger
)

func Register() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
}

func StartServer(addr string) {
	logger.Printf("Server was started on %s\n", addr)
}

func StopServer() {
	logger.Println("Server was stopped")
}

func PlayerConnected(playerAddr string) {
	logger.Printf("Player connected: %s\n", playerAddr)
}

func PlayerDisconnected(playerAddr string) {
	logger.Printf("Player disconnected: %s\n", playerAddr)
}
