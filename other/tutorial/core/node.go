package core

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type MasterNode struct {
	api     *gin.Engine
	ln      net.Listener
	svr     *grpc.Server
	nodeSvr *NodeServiceGrpcServer
}

func (node *MasterNode) Init() (err error) {
	// TODO: Move application port to .env or whateva is a Golang alternative
	node.ln, err = net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	node.svr = grpc.NewServer()
	node.nodeSvr = GetNodeServiceGrpcServer()

	RegisterNodeServiceServer(node.svr, node.nodeSvr)

	node.api = gin.Default()
	node.api.POST("/tasks", func(c *gin.Context) {
		var payload struct {
			Cmd string `json:"cmd"`
		}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		node.nodeSvr.CmdChannel <- payload.Cmd
		c.AbortWithStatus(http.StatusOK)
	})
	return nil
}

func (node *MasterNode) Start() {
	go node.svr.Serve(node.ln)

	// TODO: Same as before
	_ = node.api.Run(":9092")

	node.svr.Stop()
}

var node *MasterNode

// GetMasterNode Returns the node
// TODO: -_-

func GetMasterNode() *MasterNode {
	if node == nil {
		node = &MasterNode{}

		if err := node.Init(); err != nil {
			panic(err)
		}
	}
	return node
}
