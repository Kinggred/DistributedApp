package components

import (
	cst "tic-tac-toe/core/gui/layouts"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GetLobbyRight() *fyne.Container {

	refreshButton := widget.NewButton("Refresh", func() {})
	joinButton := widget.NewButton("Join", func() {})

	lobbyEntryContainer := GetLobbyEntryContainer()
	lobbyContainer := container.NewVBox(lobbyEntryContainer, spacer)
	scrollableLobbyListContainer := container.NewVScroll(lobbyContainer)
	scrollableLobbyListContainer.SetMinSize(fyne.NewSize(scrollableLobbyListContainer.Size().Width, 50))

	buttonRow := container.New(
		cst.NewRatioHLayout(0.685, 0.01, 0.285),
		joinButton,
		spacer,
		refreshButton,
	)

	onlyOpenCheckbox := widget.NewCheck("Only open", func(b bool) {})
	onlyLocalCheckbox := widget.NewCheck("Only local", func(b bool) {})

	checkboxRow := container.New(
		cst.NewRatioHLayout(0.45, 0.1, 0.45),
		onlyOpenCheckbox,
		spacer,
		onlyLocalCheckbox,
	)

	rightContainer := container.New(
		layout.NewVBoxLayout(),
		scrollableLobbyListContainer,
		spacer,
		buttonRow,
		checkboxRow)

	return rightContainer
}
