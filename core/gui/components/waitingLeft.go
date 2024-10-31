package components

import (
	"image/color"

	glo "tic-tac-toe/core/global"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func GetWaitingLeft() *fyne.Container {
	userNameText := canvas.NewText(glo.LocalPlayer.GetUsername(), color.White)
	userNameText.Alignment = fyne.TextAlignCenter
	userNameText.TextSize = 24

	waitingLeftContainer := container.NewVBox(
		userNameText,
		spacer,
		GetLobbyViewContainer(),
	)
	return waitingLeftContainer
}
