package common

import "sync"

type LocalLobby struct {
	Lobby *Lobby
	Mu    sync.RWMutex
}

type Lobby struct {
	ID        string
	Name      string
	Owner     Player
	IsFull    bool
	IsRunning bool
	Players   Players
}

type Lobbies struct {
	Lobbies map[string]*Lobby
	Mu      sync.RWMutex
}

func NewLobbies() Lobbies {
	return Lobbies{
		Lobbies: make(map[string]*Lobby),
	}
}

func (l *Lobbies) GetLobbies() []*Lobby {
	l.Mu.RLock()
	defer l.Mu.RUnlock()

	lobbies := make([]*Lobby, 0, len(l.Lobbies))
	for _, lobby := range l.Lobbies {
		lobbies = append(lobbies, lobby)
	}
	return lobbies
}

func (l *Lobbies) SetLobbies(lobbies []*Lobby) {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	newLobbies := make(map[string]*Lobby)
	for _, lobby := range newLobbies {
		l.Lobbies[lobby.ID] = lobby
	}
}

func (l *Lobbies) AppendOrUpdateLobby(lobby *Lobby) {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	l.Lobbies[lobby.ID] = lobby
}

func (l *Lobbies) RemoveLobbies(id string) {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	delete(l.Lobbies, id)

}

func (ll *LocalLobby) GetLobby() *Lobby {
	ll.Mu.RLock()
	defer ll.Mu.Unlock()
	return ll.Lobby
}

func (ll *LocalLobby) SetLobby(lobby *Lobby) {
	ll.Mu.Lock()
	defer ll.Mu.Unlock()
	ll.Lobby = lobby
}
