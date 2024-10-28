package components

import (
	"image/color"
	clr "tic-tac-toe/core/gui/colors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func GetPlayerContainer() *fyne.Container {
	// TODO: Replace hardcoded values with some king of logic
	playerNameText := canvas.NewText("Player Name", clr.Border)

	playerContainerPadded := container.NewPadded(playerNameText)

	// Border might be overrated
	border := canvas.NewRectangle(color.Transparent)
	border.StrokeWidth = 1
	border.StrokeColor = clr.Border

	playerContainer := container.NewStack(
		border,
		playerContainerPadded,
	)

	return playerContainer
}
