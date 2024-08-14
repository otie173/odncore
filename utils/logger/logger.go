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

func PlayerConnected(playerID byte) {
	logger.Printf("Player connected: id%d\n", playerID)
}

func PlayerDisconnected(playerID byte) {
	logger.Printf("Player disconnected: id%d\n", playerID)
}
