package networking

import (
	"context"
	"net"
	com "tic-tac-toe/core/common"
	conf "tic-tac-toe/core/config"
	glo "tic-tac-toe/core/global"
	"time"

	"github.com/taskcluster/slugid-go/slugid"
	"google.golang.org/grpc"
)

type LobbyServer struct {
	UnimplementedLobbySearchServiceServer
	lobby *com.Lobby
}

func (server *LobbyServer) mapPlayer(player *com.Player) *Player {
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
			Id:        server.lobby.ID,
			Owner:     server.mapPlayer(&glo.PlayerInfo),
			IsFull:    server.lobby.IsFull,
			IsRunning: server.lobby.IsRunning,
		}
		server.lobby.Mu.RUnlock()

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
	server.lobby.Mu.Lock()
	defer server.lobby.Mu.Unlock()
	server.lobby.Owner = &glo.PlayerInfo
	return &RegisterOwnerResponse{
		Accepted: true,
	}, nil
}

func RunServer(name string) {
	lobby := &com.Lobby{
		ID:        slugid.Nice(),
		Name:      name,
		Owner:     &glo.PlayerInfo,
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
	conf.ServerLogger.Printf("LobbyServer ('" + name + "') is running on port 50051")

	if err := grpcServer.Serve(listener); err != nil {
		conf.ServerLogger.Fatalf("Done shat myself")
	}
}
