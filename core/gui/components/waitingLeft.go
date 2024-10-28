package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func GetWaitingLeft() *fyne.Container {
	userNameText := canvas.NewText("Username", color.White)

	waitingLeftContainer := container.NewVBox(
		userNameText,
		spacer,
		GetLobbyViewContainer(),
	)
	return waitingLeftContainer
}
