package common

import (
	"sync"

	"github.com/taskcluster/slugid-go/slugid"
)

type Player struct {
	ID       string
	Username string
	Shape    int8
	Mu       sync.RWMutex
}

type Players struct {
	Host       Player
	Guest      Player
	Spectators []Player
}

func NewPlayer() Player {
	id := slugid.Nice()
	return Player{
		ID: id,
		// Defaults to id to avoid shit
		Username: id,
		// Defaults to empty field
		Shape: 0,
	}
}
