package main

import (
	"context"
	"sync"
	config "tic-tac-toe/core/config"
	networking "tic-tac-toe/core/networking"
	"time"
)

func main() {
	config.LoadVariables()
	config.SetupLoggers()

	log := config.AppLogger

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(2)

	// TODO: Start server on command of a user
	go func() {
		defer wg.Done()
		log.Printf("Trying to start the Lobby")
		networking.RunServer()
	}()

	if config.CONFIG.SERVER_ONLY == false {
		go func() {
			defer wg.Done()
			log.Printf("Trying to start a client")
			networking.RunClient(ctx)
		}()
	} else {
		defer wg.Done()
	}

	// Test
	time.Sleep(60 * time.Second)
	cancel()

	wg.Wait()

}
