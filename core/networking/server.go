package networking

import (
	"context"
	"net"
	"sync"
	conf "tic-tac-toe/core/config"
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
			conf.ServerLogger.Printf("Client broke the connection")
			return nil
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
		conf.ServerLogger.Fatal("Failed to listen: ", err)
	}

	grpcServer := grpc.NewServer()
	lobbyServer := &LobbyServer{
		lobby: lobby,
	}
	RegisterLobbySearchServiceServer(
		grpcServer,
		lobbyServer,
	)
	conf.ServerLogger.Printf("LobbyServer is running on port 50051")

	if err := grpcServer.Serve(listener); err != nil {
		conf.ServerLogger.Fatalf("Done shat myself")
	}
}
