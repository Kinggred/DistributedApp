package interactions

import (
	"sync"
	log "tic-tac-toe/core/config/logging"
	glo "tic-tac-toe/core/global"
	netw "tic-tac-toe/core/networking"

	"github.com/sirupsen/logrus"
)

func HostAGame(name string) {
	var wg sync.WaitGroup
	log := log.AppLogger.WithFields(logrus.Fields{"name": name})
	if !glo.IsServerRunning() {
		go func() {
			defer wg.Done()
			log.Printf("Trying to start the Lobby ")
			glo.ChangeServerStatus()
			netw.RunServer(name)
		}()
	}
}
