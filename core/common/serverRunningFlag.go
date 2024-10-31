package common

import "sync"

type ServerStatusStruct struct {
	ServerRunning bool
	Mu            sync.RWMutex
}
