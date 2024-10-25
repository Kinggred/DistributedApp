package main

import (
	"log"
	"sync"
	config "tic-tac-toe/core/config"
	networking "tic-tac-toe/core/networking"
)

func main() {

	config.LoadVariables()
	var wg sync.WaitGroup
	wg.Add(2)

	// TODO: Start server on command of a user
	go func() {
		defer wg.Done()
		log.SetPrefix("Server: ")
		log.Printf("Trying to start the Lobby")
		networking.RunServer()
	}()

	if config.CONFIG.SERVER_ONLY == false {
		go func() {
			defer wg.Done()
			log.SetPrefix("Client: ")
			log.Printf("Trying to start a client")
			networking.RunClient()
		}()
	} else {
		defer wg.Done()
	}

	wg.Wait()
}
