package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	glo "tic-tac-toe/core/global"
	cst "tic-tac-toe/core/gui/layouts"
)

func GetGameScreen() *fyne.Container {

	pussyButton := widget.NewButton("Surrender", func() {})

	playersContainer := container.NewVBox(
		spacer,
		GetPlayerContainer(glo.LocalPlayer.GetUsername()),
		GetPlayerContainer("Placeholder"),
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
