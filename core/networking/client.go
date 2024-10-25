package networking

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
)

const (
	// TODO: This has to be set by the user
	defaultName = "nick"

	// TODO: Move out to some kind of .env
	ApplicationPort = "50051"
	VersionMain     = 0
	VersionMin      = 1
)

func seekActiveHost(address string, port string) {
	conn, err := grpc.Dial(address+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect to server %v", err)
	}

	client := NewLobbySearchServiceClient(conn)
	in := &LobbyRequest{
		VersionMain: VersionMain,
		VersionMin:  VersionMin,
	}
	stream, err := client.RequestActiveLobbies(context.Background(), in)
	if err != nil {
		log.Fatalf("Error during openning stream")
	}

	done := make(chan bool)

	go func() {
		log.SetPrefix("Client: ")
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				log.Fatalf("Cannot receive %v", err)
			}
			// TODO: Implement some LOOOOOOGIC
			log.Printf("Response received: %s", response)
		}
	}()

	<-done
	log.Printf("Finished")
}

func RunClient() {
	// TODO: This seeks localhost only for now.
	// Some kind of Goroutines for multiple connections might be required
	// Update: Looks like goroutines and channels will do the trick
	log.Printf("Started a client")
	seekActiveHost("", ApplicationPort)
}
