package core

import (
	"context"
)

type NodeServiceGrpcServer struct {
	UnimplementedNodeServiceServer

	// Channel to recieve command
	CmdChannel chan string
}

func (node NodeServiceGrpcServer) ReportStatus(ctx context.Context, request *Request) (*Response, error) {
	return &Response{Data: "ok"}, nil
}

func (node NodeServiceGrpcServer) AssignTask(request *Request, server NodeService_AssignTaskServer) error {
	for {
		select {
		case cmd := <-node.CmdChannel:
			// Recieve command and send to worker node (client)
			if err := server.Send(&Response{Data: cmd}); err != nil {
				return err
			}
		}
	}
}

var server *NodeServiceGrpcServer

// GetNodeServiceGrpcServer singleton service ( TODO: Figure Out TF Happens Here )
func GetNodeServiceGrpcServer() *NodeServiceGrpcServer {
	if server == nil {
		server = &NodeServiceGrpcServer{
			CmdChannel: make(chan string),
		}
	}
	return server
}
