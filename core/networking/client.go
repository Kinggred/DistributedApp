package networking

import (
	"context"
	"io"
	"log"
	"sync"

	conf "tic-tac-toe/core/config"

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
	// TODO: Fix, WithInsecure is obsolete
	// TODO: I forgot why I made this, but something else has to be done for sure
	conn, err := grpc.Dial(address+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect to server %v", err)
	}

	client := NewLobbySearchServiceClient(conn)
	in := &LobbyRequest{
		VersionMain: int32(conf.CONFIG.VERSION_MAIN),
		VersionMin:  int32(conf.CONFIG.VERSION_MAIN),
	}
	stream, err := client.RequestActiveLobbies(context.Background(), in)
	if err != nil {
		log.Fatalf("Error during openning stream")
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			log.Printf("Finished")
		}
		if err != nil {
			log.Fatalf("Cannot receive %v", err)
		}
		// TODO: Implement some LOOOOOOGIC
		log.Printf("Response received: %s", response.Id)
	}

}

func RunClient() {

	// TODO: Implement channels so We have a way to access lobby status from outside
	// TODO: This has to be updatable prefferabbly without reccontcting already established links.
	// No clue how to do this, at least for now.
	// REMINDER: Due to the nature of this kerfuffle I have to setup separate machines ( docker possibly? )
	// to be able to test it.
	log.Printf("Started a client")
	var wg sync.WaitGroup
	activeHosts := ScanNetwork()

	wg.Add(len(activeHosts))
	for _, ip := range activeHosts {
		go func() {
			defer wg.Done()
			seekActiveHost(ip, ApplicationPort)
		}()
	}
	wg.Wait()
}
