package main

import (
	"fmt"
	"tic-tac-toe/core"
)

func main() {
	fmt.Print("Enter node type (master / worker): ")
	var nodeType string
	fmt.Scanln(&nodeType)

	switch nodeType {
	case "master":
		core.GetMasterNode().Start()
	case "worker":
		core.GetWorkerNode().Start()
	default:
		panic("Invalid node request: " + nodeType)
	}
}
