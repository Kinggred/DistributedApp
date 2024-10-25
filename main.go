package main

import (
	"log"
	"sync"
	networking "tic-tac-toe/core/networking"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// TODO: Start server on command of a user
	go func() {
		defer wg.Done()
		log.SetPrefix("Server: ")
		log.Printf("Trying to start the Lobby")
		networking.RunServer()
	}()

	go func() {
		defer wg.Done()
		log.SetPrefix("Client: ")
		log.Printf("Trying to start a client")
		networking.RunClient()
	}()

	wg.Wait()
}
