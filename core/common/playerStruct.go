package common

import (
	"sync"

	"github.com/taskcluster/slugid-go/slugid"
)

type Player struct {
	ID       string
	Username string
	Shape    int8
}

// After Joining or creating a lobby this data will be duplicated in Players struct.
// But we wont think about it too much.
type LocalPlayer struct {
	player Player
	mu     sync.RWMutex
}

func newPlayer() Player {
	id := slugid.Nice()
	return Player{
		ID:       id,
		Username: "",
		// Defaults to empty field
		Shape: 0,
	}
}

func NewLocalPlayer() LocalPlayer {
	return LocalPlayer{
		player: newPlayer(),
	}
}

func (lp *LocalPlayer) Get() Player {
	lp.mu.RLock()
	defer lp.mu.RUnlock()
	return lp.player
}

func (lp *LocalPlayer) GetUsername() string {
	lp.mu.RLock()
	defer lp.mu.RUnlock()
	return lp.player.Username
}

func (lp *LocalPlayer) SetUsername(username string) {
	lp.mu.Lock()
	defer lp.mu.Unlock()
	lp.player.Username = username
}

func (lp *LocalPlayer) GetID() string {
	lp.mu.RLock()
	defer lp.mu.RUnlock()
	return lp.player.ID
}

func (lp *LocalPlayer) GetShape() int8 {
	lp.mu.RLock()
	defer lp.mu.RUnlock()
	return lp.player.Shape
}

func (lp *LocalPlayer) SetShape(shape int8) {
	lp.mu.Lock()
	defer lp.mu.Unlock()
	lp.player.Shape = shape
}
