package networking

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"github.com/taskcluster/slugid-go/slugid"
	"google.golang.org/grpc"
)

type Lobby struct {
	ID        string
	Owner     string
	IsFull    bool
	IsRunning bool
	mu        sync.RWMutex
}

type LobbyServer struct {
	UnimplementedLobbySearchServiceServer
	lobby *Lobby
}

func (server *LobbyServer) RequestActiveLobbies(request *LobbyRequest, stream LobbySearchService_RequestActiveLobbiesServer) error {
	for {
		server.lobby.mu.RLock()
		proposal := &LobbyProposal{
			Id:        server.lobby.ID,
			Owner:     server.lobby.Owner,
			IsFull:    server.lobby.IsFull,
			IsRunning: server.lobby.IsRunning,
		}
		server.lobby.mu.RUnlock()

		if err := stream.Send(proposal); err != nil {
			// CONSIDER: This could fail silently
			log.Fatalf("Error during response")
		}

		// Possible adjustments
		time.Sleep(10 * time.Second)
	}
}

func (server *LobbyServer) RequestLobbyJoin(context context.Context, request *JoinLobbyRequest) (*JoinLobbyResponse, error) {
	return &JoinLobbyResponse{
		Id:  server.lobby.ID,
		Key: "some_key", // Logic to generate key should be added
	}, nil
}

func (server *LobbyServer) RegisterOwnerRequest(context context.Context, request *RegisterOwnerRequest) (*RegisterOwnerResponse, error) {
	server.lobby.mu.Lock()
	defer server.lobby.mu.Unlock()
	server.lobby.Owner = request.Id
	return &RegisterOwnerResponse{
		Accepted: true,
	}, nil
}

func RunServer() {
	lobby := &Lobby{
		ID:        slugid.Nice(),
		Owner:     "",
		IsFull:    false,
		IsRunning: false,
	}

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	grpcServer := grpc.NewServer()
	lobbyServer := &LobbyServer{
		lobby: lobby,
	}
	RegisterLobbySearchServiceServer(
		grpcServer,
		lobbyServer,
	)
	log.Println("LobbyServer is running on port 50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Done shat myself")
	}
}
