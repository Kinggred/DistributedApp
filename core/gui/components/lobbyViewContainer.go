package components

import (
	cst "tic-tac-toe/core/gui/layouts"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	glo "tic-tac-toe/core/global"
)

func GetLobbyViewContainer() *fyne.Container {

	spectatorContainer := container.New(
		cst.NewRatioHLayout(0.05, 0.9, 0.05),
		spacer,
		GetPlayerContainer("Placeholder"),
		spacer,
	)

	spectatorsContainer := container.NewVBox(
		spectatorContainer,
		spacer,
	)
	spectatorsScrollableContainer := container.NewVScroll(spectatorsContainer)

	startButton := widget.NewButton("Start", func() {})

	viewContainer := container.NewVBox(
		GetPlayerContainer(glo.LocalPlayer.GetUsername()),
		GetPlayerContainer("Placeholder"),
		spectatorsScrollableContainer,
		startButton,
	)

	return viewContainer
}
