package common

import (
	"sync"
)

type Players struct {
	Host       Player
	Guest      Player
	Spectators []Player
	Mu         sync.RWMutex
}

func (p *Players) NewPlayers() Players {
	return Players{
		Host:       newPlayer(),
		Spectators: []Player{},
	}
}

func (p *Players) GetHost() *Player {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return &p.Host
}

func (p *Players) SetHost(host Player) {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.Host = host
}

func (p *Players) GetGuest() *Player {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return &p.Guest
}

func (p *Players) SetGuest(guest Player) {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.Guest = guest
}

func (p *Players) GetSpectators() *[]Player {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return &p.Spectators
}

func (p *Players) SetSpectators(spectators []Player) {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.Spectators = spectators
}

func (p *Players) AppendSpectators(spectator Player) {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.Spectators = append(p.Spectators, spectator)
}

func (p *Players) RemoveSpectator(spectator Player) {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	lastIndex := len(p.Spectators)
	for i, user := range p.Spectators {
		if user == spectator {
			p.Spectators[i] = p.Spectators[lastIndex]
			p.Spectators = p.Spectators[:lastIndex]
		}
	}
}
