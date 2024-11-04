package networking

import (
	"context"
	"net"
	com "tic-tac-toe/core/common"
	log "tic-tac-toe/core/config/logging"
	glo "tic-tac-toe/core/global"
	"time"

	"github.com/taskcluster/slugid-go/slugid"
	"google.golang.org/grpc"
)

type LobbyServer struct {
	UnimplementedLobbySearchServiceServer
	lobby *com.LocalLobby
}

func (server *LobbyServer) mapPlayer(player com.Player) *Player {
	return &Player{
		ID:       player.ID,
		Username: player.Username,
		Shape:    int32(player.Shape),
	}
}

func (server *LobbyServer) RequestActiveLobbies(request *LobbyRequest, stream LobbySearchService_RequestActiveLobbiesServer) error {
	for {
		server.lobby.Mu.RLock()
		proposal := &LobbyProposal{
			Id:        server.lobby.Lobby.ID,
			Name:      server.lobby.Lobby.Name,
			Owner:     server.mapPlayer(glo.LocalPlayer.Get()),
			IsFull:    server.lobby.Lobby.IsFull,
			IsRunning: server.lobby.Lobby.IsRunning,
		}
		server.lobby.Mu.RUnlock()

		if err := stream.Send(proposal); err != nil {
			log.ServerLogger.Printf("Client broke the connection")
			return nil
		}

		// Possible adjustments
		time.Sleep(10 * time.Second)
	}
}

func (server *LobbyServer) RequestLobbyJoin(context context.Context, request *JoinLobbyRequest) (*JoinLobbyResponse, error) {
	server.lobby.Mu.Lock()
	defer server.lobby.Mu.Unlock()
	return &JoinLobbyResponse{
		Id:  server.lobby.Lobby.ID,
		Key: "some_key", // Logic to generate key should be added
	}, nil
}

func (server *LobbyServer) RegisterOwnerRequest(context context.Context, request *RegisterOwnerRequest) (*RegisterOwnerResponse, error) {
	server.lobby.Mu.Lock()
	defer server.lobby.Mu.Unlock()
	server.lobby.Lobby.Owner = glo.LocalPlayer.Get()
	return &RegisterOwnerResponse{
		Accepted: true,
	}, nil
}

func RunServer(name string) {
	lobby := &com.Lobby{
		ID:        slugid.Nice(),
		Name:      name,
		Owner:     glo.LocalPlayer.Get(),
		IsFull:    false,
		IsRunning: false,
	}

	glo.LocalLobbyData.SetLobby(lobby)
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.ServerLogger.Fatal("Failed to listen: ", err)
	}

	grpcServer := grpc.NewServer()
	lobbyServer := &LobbyServer{
		lobby: &glo.LocalLobbyData,
	}
	RegisterLobbySearchServiceServer(
		grpcServer,
		lobbyServer,
	)
	log.ServerLogger.Printf("LobbyServer ('" + name + "') is running on port 50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.ServerLogger.Fatalf("Done shat myself")
	}
}
