package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	cst "tic-tac-toe/core/gui/layouts"
)

func GetGameGrid() *fyne.Container {

	// TODO: This will have to be eddited to respond to changes in some Global
	// Map containing current game status.
	var gameTiles []fyne.CanvasObject

	for i := 0; i < 9; i++ {
		tile := canvas.NewRectangle(color.White)
		tile.StrokeColor = color.Black
		tile.StrokeWidth = 4
		tile.FillColor = color.Transparent

		button := widget.NewButton("", func() {})
		cont := container.NewStack(button, tile)
		gameTiles = append(gameTiles, cont)
	}

	squaredGameGrid := container.New(
		cst.NewSquareLayout(237),
		gameTiles...,
	)

	return squaredGameGrid
}
