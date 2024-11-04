package gameLogic

import (
	"context"
	com "tic-tac-toe/core/common"
	glo "tic-tac-toe/core/global"
	gui "tic-tac-toe/core/gui"

	"fyne.io/fyne/v2/app"
)

type Game struct {
	gameFields  com.FieldStatus
	playerShape int8
	ServerId    string
}

func GameLoop() {
	glo.InitGUIStateGlobal()
	glo.InitLobbiesData()
	glo.InitPlayersGlobal()
	glo.InitLocalPlayerGlobal()

	ctx, killClient := context.WithCancel(context.Background())
	glo.InitClientKillerGlobal(killClient)

	app := app.New()

	window := gui.RenderStartingScreen(app, ctx)

	window.ShowAndRun()
}
