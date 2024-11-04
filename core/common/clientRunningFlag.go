package common

import (
	"context"
	"sync"
)

type ClientKiller struct {
	kill            context.CancelFunc
	isClientRunning bool
	mu              sync.RWMutex
}

func NewClientKiller(cf context.CancelFunc) ClientKiller {
	return ClientKiller{
		kill:            cf,
		isClientRunning: false,
	}
}

func (ck *ClientKiller) SetKillFunction(function context.CancelFunc) {
	ck.mu.Lock()
	defer ck.mu.Unlock()
	ck.kill = function
}

func (ck *ClientKiller) Kill() {
	ck.mu.Lock()
	defer ck.mu.Unlock()
	if ck.isClientRunning {
		ck.kill()
		ck.kill = nil
	}
}

func (ck *ClientKiller) GetStatus() bool {
	ck.mu.RLock()
	defer ck.mu.RUnlock()
	return ck.isClientRunning
}

func (ck *ClientKiller) RegisterClientStart() {
	ck.mu.Lock()
	defer ck.mu.Unlock()
	ck.isClientRunning = true
}

func (ck *ClientKiller) RegisterClientShutdown() {
	ck.mu.Lock()
	defer ck.mu.Unlock()
	ck.isClientRunning = false
}
