package components

import (
	"context"
	cst "tic-tac-toe/core/gui/layouts"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	glo "tic-tac-toe/core/global"
)

func GetLobbyViewContainer(ctx context.Context) *fyne.Container {

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
	abordButton := widget.NewButton("Abort", func() {
		// TODO: Server stopping logic
		glo.GUIState.Mu.Lock()
		defer glo.GUIState.Mu.Unlock()
		glo.GUIState.LeftContainer.Objects = GetLobbyLeft(ctx).Objects
		glo.GUIState.RightContainer.Objects = GetLobbyRight(ctx).Objects
	})

	buttonsContainer := container.New(cst.NewRatioHLayout(0.685, 0.030, 0.385), startButton, spacer, abordButton)

	viewContainer := container.NewVBox(
		GetPlayerContainer(glo.LocalPlayer.GetUsername()),
		GetPlayerContainer("Placeholder"),
		spectatorsScrollableContainer,
		buttonsContainer,
	)

	return viewContainer
}
