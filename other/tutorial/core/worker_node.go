package core

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"google.golang.org/grpc"
)

type WorkerNode struct {
	conn *grpc.ClientConn
	c    NodeServiceClient
}

func (node *WorkerNode) Init() (err error) {
	node.conn, err = grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return err
	}

	// Why does this fuck itself
	node.c = NewNodeServiceClient(node.conn)

	return nil
}

func (node *WorkerNode) Start() {
	fmt.Println("Worker node started")

	_, _ = node.c.ReportStatus(context.Background(), &Request{})

	stream, _ := node.c.AssignTask(context.Background(), &Request{})
	for {
		res, err := stream.Recv()

		if err != nil {
			return
		}

		fmt.Println("received command: ", res.Data)

		parts := strings.Split(res.Data, " ")
		if err := exec.Command(parts[0], parts[1:]...).Run(); err != nil {
			fmt.Println(err)
		}

	}
}

var workerNode *WorkerNode

func GetWorkerNode() *WorkerNode {
	if workerNode == nil {
		workerNode = &WorkerNode{}

		if err := workerNode.Init(); err != nil {
			panic(err)
		}
	}

	return workerNode
}
