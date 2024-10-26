package networking

import (
	"context"
	"io"
	"sync"

	conf "tic-tac-toe/core/config"

	"google.golang.org/grpc"
)

const (
	// TODO: This has to be set by the user
	defaultName = "nick"

	ApplicationPort = "50051"
)

func seekActiveHost(ctx context.Context, address string, port string) {
	// TODO: Fix, WithInsecure is obsolete
	// TODO: I forgot why I made this, but something else has to be done for sure

	conn, err := grpc.Dial(address+":"+port, grpc.WithInsecure())
	if err != nil {
		conf.ClientLogger.Fatalf("Cannot connect to server %v", err)
	}

	client := NewLobbySearchServiceClient(conn)
	in := &LobbyRequest{
		VersionMain: int32(conf.CONFIG.VERSION_MAIN),
		VersionMin:  int32(conf.CONFIG.VERSION_MAIN),
	}
	stream, err := client.RequestActiveLobbies(context.Background(), in)
	if err != nil {
		conf.ClientLogger.Fatalf("Error during openning stream")
	}

	for {
		select {
		case <-ctx.Done():
			conf.ClientLogger.Printf("Stopping connection")
			conn.Close()
			return
		default:
			response, err := stream.Recv()
			if err == io.EOF {
				conf.ClientLogger.Printf("Finished")
			}
			if err != nil {
				conf.ClientLogger.Fatalf("Cannot receive %v", err)
			}
			// TODO: Implement some LOOOOOOGIC
			conf.ClientLogger.Printf("Response received: %s", response.Id)
		}
	}
}

func RunClient(ctx context.Context) {

	// TODO: Implement channels so We have a way to access lobby status from outside
	// TODO: This has to be updatable prefferabbly without reconnecting already established links.
	// No clue how to do this, at least for now.
	// Changed my mind, reconnecting will do.
	conf.ClientLogger.Printf("Started a client")
	var wg sync.WaitGroup
	activeHosts := ScanNetwork()

	wg.Add(len(activeHosts))
	for _, ip := range activeHosts {
		go func() {
			defer wg.Done()
			seekActiveHost(ctx, ip, ApplicationPort)
		}()
	}
	wg.Wait()
}
