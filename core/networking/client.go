package networking

import (
	"context"
	"io"
	"sync"

	com "tic-tac-toe/core/common"
	conf "tic-tac-toe/core/config"
	log "tic-tac-toe/core/config/logging"
	glo "tic-tac-toe/core/global"

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
	log := log.ClientLogger.WithField("ip", address).WithField("port", port)

	conn, err := grpc.Dial(address+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Errorf("Cannot connect to server %v", err)
		return
	}
	defer conn.Close()

	client := NewLobbySearchServiceClient(conn)
	in := &LobbyRequest{
		VersionMain: int32(conf.CONFIG.VERSION_MAIN),
		VersionMin:  int32(conf.CONFIG.VERSION_MAIN),
	}

	dataChannel := make(chan *LobbyProposal)

	stream, err := client.RequestActiveLobbies(context.Background(), in)
	defer stream.CloseSend()

	if err != nil {
		log.Errorf("Error during openning stream")
		return
	}

	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				log.Printf("Finished")
				return
			}
			if err != nil {
				log.Errorf("Cannot receive %v", err)
				return
			}
			dataChannel <- response
		}
	}()

	for {
		select {
		case <-ctx.Done():
			conn.Close()
			stream.CloseSend()
			log.Printf("Stopping connection")
			return
		case response := <-dataChannel:
			log.Printf("Response received: %s", response.Id)
			owner := com.Player{
				ID:       response.GetOwner().ID,
				Username: response.GetOwner().Username,
				Shape:    int8(response.GetOwner().Shape),
			}
			lobby := com.Lobby{
				ID:        response.GetId(),
				Name:      response.GetName(),
				Owner:     owner,
				IsFull:    response.GetIsFull(),
				IsRunning: response.GetIsRunning(),
			}

			glo.LobbiesData.AppendOrUpdateLobby(&lobby)
		}
	}
}

func RunClient(ctx context.Context) {

	// TODO: Implement channels so We have a way to access lobby status from outside
	// TODO: This has to be updatable prefferabbly without reconnecting already established links.
	// No clue how to do this, at least for now.
	// Changed my mind, reconnecting will do.
	log.ClientLogger.Printf("Started a client")
	glo.KillClient.RegisterClientStart()
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
	glo.KillClient.RegisterClientShutdown()
}
