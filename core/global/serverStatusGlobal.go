package gloabals

import (
	com "tic-tac-toe/core/common"
)

var serverStatus com.ServerStatusStruct

func initServerStatusGlobal() {
	serverStatus = com.ServerStatusStruct{
		ServerRunning: false,
	}
}

func IsServerRunning() bool {
	serverStatus.Mu.RLock()
	defer serverStatus.Mu.RUnlock()
	return serverStatus.ServerRunning
}

func ChangeServerStatus() {
	serverStatus.Mu.Lock()
	defer serverStatus.Mu.Unlock()
	current := serverStatus.ServerRunning
	serverStatus.ServerRunning = !current
}
