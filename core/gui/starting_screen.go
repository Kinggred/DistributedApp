package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	cmp "tic-tac-toe/core/gui/components"
	cst "tic-tac-toe/core/gui/layouts"
)

func RenderStartingScreen() {
	app := app.New()
	window := app.NewWindow("Tic-Tac-Toe")

	spacer := layout.NewSpacer()

	leftContainer := cmp.GetLobbyLeft()
	rightContainer := cmp.GetLobbyRight()
	toolbarContainer := cmp.GetToolbar()

	mainContainer := container.New(cst.NewRatioHLayout(0.45, 0.1, 0.45), leftContainer, spacer, rightContainer)

	combinedContainer := container.New(cst.NewRatioVLayout(0.944, 0.005, 0.06), mainContainer, spacer, toolbarContainer)
	paddedContainer := container.NewPadded(combinedContainer)

	window.SetContent(paddedContainer)
	window.Resize(fyne.NewSize(550, 300))
	window.ShowAndRun()
}
