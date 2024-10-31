package interactions

import (
	"sync"
	cfg "tic-tac-toe/core/config"
	glo "tic-tac-toe/core/global"
	netw "tic-tac-toe/core/networking"
)

func HostAGame(name string) {
	var wg sync.WaitGroup

	if !glo.IsServerRunning() {
		go func() {
			defer wg.Done()
			cfg.AppLogger.Printf("Trying to start the Lobby")
			glo.ChangeServerStatus()
			netw.RunServer(name)
		}()
	}
}
