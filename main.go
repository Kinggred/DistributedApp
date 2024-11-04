package main

import (
	"flag"
	"tic-tac-toe/core/config"
	loggers "tic-tac-toe/core/config/logging"
	game "tic-tac-toe/core/gameLogic"
	networking "tic-tac-toe/core/networking"
)

func main() {
	config.LoadVariables()
	loggers.SetupLoggers()
	log := loggers.AppLogger

	mode := flag.String("mode", "", "Set to 'server_only' if no gui or client is desired")
	flag.Parse()

	switch *mode {
	case "server_only":
		log.Printf("Starting in server_only mode")
		log.Printf("Trying to start the Lobby")
		networking.RunServer("")
	default:
		game.GameLoop()
	}

}
