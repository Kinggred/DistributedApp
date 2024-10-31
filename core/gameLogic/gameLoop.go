package gameLogic

import (
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
	glo.InitPlayersGlobal()
	glo.InitLocalPlayerGlobal()

	app := app.New()
	window := gui.RenderStartingScreen(app)

	window.ShowAndRun()
}
