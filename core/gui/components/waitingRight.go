package components

import (
	"image/color"

	glo "tic-tac-toe/core/global"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func GetWaitingRight() *fyne.Container {
	userNameText := canvas.NewText(glo.LocalPlayer.GetUsername(), color.White)
	userNameText.Alignment = fyne.TextAlignCenter
	userNameText.TextSize = 24

	chatText := canvas.NewText("Chat placeholder", color.White)
	chatText.Alignment = fyne.TextAlignCenter
	chatText.TextSize = 24

	waitingRightContainer := container.NewVBox(
		userNameText,
		spacer,
		chatText,
	)
	return waitingRightContainer
}
