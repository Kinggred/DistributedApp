package common

import "sync"

type Lobby struct {
	ID        string
	Name      string
	Owner     Player
	IsFull    bool
	IsRunning bool
	Players   Players
	Mu        sync.RWMutex
}
