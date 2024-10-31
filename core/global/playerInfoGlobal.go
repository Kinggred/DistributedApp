package gloabals

import (
	com "tic-tac-toe/core/common"
)

var PlayerInfo com.Player

func InitPlayerGlobal() {
	PlayerInfo = com.NewPlayer()
}
