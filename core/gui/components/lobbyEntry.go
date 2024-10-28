package components

import (
	"image/color"
	clr "tic-tac-toe/core/gui/colors"
	cst "tic-tac-toe/core/gui/layouts"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GetLobbyEntryContainer() *fyne.Container {
	// TODO: Replace hardcoded values with some king of logic
	lobbyNameText := canvas.NewText("Lobby Name", clr.Green)
	playerCountText := canvas.NewText("1/2", clr.Green)

	lobbyContainerContents := container.New(
		cst.NewRatioHLayout(0.15, 0.685, 0.02, 0.270),
		spacer,
		lobbyNameText,
		spacer,
		playerCountText,
	)

	lobbyContainerContentsPadded := container.NewPadded(lobbyContainerContents)

	// Border might be overrated
	border := canvas.NewRectangle(color.Transparent)
	border.StrokeWidth = 1
	border.StrokeColor = clr.Green

	hiddenButton := widget.NewButton("", func() {})

	lobbyContainer := container.NewStack(
		hiddenButton,
		border,
		lobbyContainerContentsPadded,
	)

	return lobbyContainer
}
