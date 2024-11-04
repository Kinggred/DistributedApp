package components

import (
	"image/color"
	"tic-tac-toe/core/common"
	glo "tic-tac-toe/core/global"
	clr "tic-tac-toe/core/gui/colors"
	cst "tic-tac-toe/core/gui/layouts"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GetLobbyEntryContainer(lobby *common.Lobby) *fyne.Container {
	playerCount := "1/2"
	if lobby.IsFull {
		playerCount = "2/2"
	}

	lobbyNameText := canvas.NewText(lobby.Name, clr.Border)
	playerCountText := canvas.NewText(playerCount, clr.Border)

	lobbyContainerContents := container.New(
		cst.NewRatioHLayout(0.15, 0.685, 0.02, 0.270),
		spacer,
		lobbyNameText,
		spacer,
		playerCountText,
	)

	lobbyContainerContentsPadded := container.NewPadded(lobbyContainerContents)

	// Border might be overrated
	border := canvas.NewRectangle(color.Transparent)
	border.StrokeWidth = 1
	border.StrokeColor = clr.Border

	if lobby.ID == glo.LobbiesData.GetChosen() {
		border.StrokeWidth = 3
	}

	hiddenButton := widget.NewButton("", func() {
		border.StrokeWidth = 3
		glo.LobbiesData.SetChosen(lobby.ID)
	})

	lobbyContainer := container.NewStack(
		hiddenButton,
		border,
		lobbyContainerContentsPadded,
	)

	return lobbyContainer
}
