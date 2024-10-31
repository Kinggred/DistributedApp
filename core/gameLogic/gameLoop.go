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
	app := app.New()
	glo.InitGUIStateGlobal()
	window := gui.RenderStartingScreen(app)

	window.ShowAndRun()
}
