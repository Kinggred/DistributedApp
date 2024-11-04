package gloabals

import (
	"context"
	com "tic-tac-toe/core/common"
)

var KillClient com.ClientKiller

func InitClientKillerGlobal(cf context.CancelFunc) {
	KillClient = com.NewClientKiller(cf)
}
