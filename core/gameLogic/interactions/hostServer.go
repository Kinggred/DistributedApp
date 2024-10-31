package interactions

import (
	"sync"
	log "tic-tac-toe/core/config/logging"
	glo "tic-tac-toe/core/global"
	netw "tic-tac-toe/core/networking"
)

func HostAGame(name string) {
	var wg sync.WaitGroup

	if !glo.IsServerRunning() {
		go func() {
			defer wg.Done()
			log.AppLogger.Printf("Trying to start the Lobby")
			glo.ChangeServerStatus()
			netw.RunServer(name)
		}()
	}
}
