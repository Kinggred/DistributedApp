package gloabals

import (
	com "tic-tac-toe/core/common"
)

var LocalPlayer com.LocalPlayer
var PlayersInfo com.Players

func InitPlayersGlobal() {
	PlayersInfo = PlayersInfo.NewPlayers()
}

func InitLocalPlayerGlobal() {
	LocalPlayer = com.NewLocalPlayer()
}
