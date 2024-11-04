package gloabals

import (
	com "tic-tac-toe/core/common"
)

var LocalLobbyData com.LocalLobby
var LobbiesData com.Lobbies

func InitLocalLobby() {
	LocalLobbyData = com.LocalLobby{
		Lobby: &com.Lobby{},
	}
}

func InitLobbiesData() {
	LobbiesData = com.Lobbies{
		Lobbies: make(map[string]*com.Lobby),
	}
}
