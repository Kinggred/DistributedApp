package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	cst "tic-tac-toe/core/gui/layouts"
)

func GetGameScreen() *fyne.Container {

	pussyButton := widget.NewButton("Surrender", func() {})

	playersContainer := container.NewVBox(
		spacer,
		GetPlayerContainer(),
		GetPlayerContainer(),
		pussyButton,
	)

	gridContainer := container.New(
		cst.NewRatioHLayout(0.02, 0.96, 0.02),
		spacer,
		GetGameGrid(),
		spacer,
	)

	mainScreenContainer := container.New(
		cst.NewRatioVLayout(0.8, 0.2),
		gridContainer,
		playersContainer,
	)

	return mainScreenContainer
}
