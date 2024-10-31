package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	glo "tic-tac-toe/core/global"
	cmp "tic-tac-toe/core/gui/components"
	cst "tic-tac-toe/core/gui/layouts"
)

func RenderStartingScreen(app fyne.App) fyne.Window {
	window := app.NewWindow("Tic-Tac-Toe")

	spacer := layout.NewSpacer()
	glo.GUIState.Mu.Lock()
	defer glo.GUIState.Mu.Unlock()
	glo.GUIState.LeftContainer = cmp.GetLobbyLeft()
	glo.GUIState.RightContainer = cmp.GetLobbyRight()
	toolbarContainer := cmp.GetToolbar(app)

	mainContainer := container.New(cst.NewRatioHLayout(0.45, 0.1, 0.45), glo.GUIState.LeftContainer, spacer, glo.GUIState.RightContainer)

	combinedContainer := container.New(cst.NewRatioVLayout(0.944, 0.005, 0.06), mainContainer, spacer, toolbarContainer)
	paddedContainer := container.NewPadded(combinedContainer)

	window.SetContent(paddedContainer)
	window.Resize(fyne.NewSize(550, 300))
	return window
}
