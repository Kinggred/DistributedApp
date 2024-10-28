package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	cst "tic-tac-toe/core/gui/layouts"
)

func GetLobbyLeft() *fyne.Container {

	userNameWidget := widget.NewEntry()
	userNameWidget.SetPlaceHolder("Username")

	spacer := layout.NewSpacer()

	lobbyNameWidget := widget.NewEntry()
	lobbyNameWidget.SetPlaceHolder("Lobby Name")

	// TODO: Implement Logic
	lobbyStartButton := widget.NewButton("Host", func() {})

	lobbyCreationContainer := container.New(
		cst.NewRatioHLayout(0.685, 0.01, 0.285),
		lobbyNameWidget,
		spacer,
		lobbyStartButton,
	)

	leftContainer := container.New(
		layout.NewVBoxLayout(),
		userNameWidget,
		spacer,
		lobbyCreationContainer,
	)

	return leftContainer
}
