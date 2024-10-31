package main

import (
	//"context"
	"flag"
	//"sync"
	config "tic-tac-toe/core/config"
	networking "tic-tac-toe/core/networking"

	//"time"

	game "tic-tac-toe/core/gameLogic"
)

func main() {
	config.LoadVariables()
	config.SetupLoggers()

	log := config.AppLogger
	// ctx, cancel := context.WithCancel(context.Background())
	// TODO: Maybe move outside one day
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

// TODO: Cleanup, left in case of a major fuckup
// And for som gitless sneek peeks
//  	log.Printf("Starting in normal mode")
//  	var wg sync.WaitGroup
//
//		wg.Add(2)
//		// TODO: Start server on command of a user
//		go func() {
//			defer wg.Done()
//			log.Printf("Trying to start the Lobby")
//			networking.RunServer()
//		}()
//
//		go func() {
//			defer wg.Done()
//			log.Printf("Trying to start a client")
//			networking.RunClient(ctx)
//		}()
//		time.Sleep(60 * time.Second)
//		cancel()
//		wg.Wait()
