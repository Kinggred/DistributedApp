package components

import (
	"context"
	ita "tic-tac-toe/core/gameLogic/interactions"
	glo "tic-tac-toe/core/global"
	cst "tic-tac-toe/core/gui/layouts"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var lobbyListContainer fyne.Container

func updateLobbyList() {
	lobbyListContainer.Objects = nil
	for _, lobby := range glo.LobbiesData.GetLobbies() {
		lobbyEntryContainer := GetLobbyEntryContainer(lobby)
		lobbyContainer := container.NewVBox(lobbyEntryContainer, spacer)
		lobbyListContainer.Objects = append(lobbyListContainer.Objects, lobbyContainer)
	}
}

func GetLobbyRight(ctx context.Context) *fyne.Container {
	keepRefreshing := true

	refreshButton := widget.NewButton("Refresh", func() {
		ita.RunClientInGoroutine(ctx)
	})

	joinButton := widget.NewButton("Join", func() {})

	lobbyListContainer = *container.NewVBox()
	scrollableLobbyListContainer := container.NewVScroll(&lobbyListContainer)
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

	// This sucks, but works
	go func() {
		for range time.Tick(3 * time.Second) {
			if !keepRefreshing {
				break
			}
			updateLobbyList()
		}
	}()

	return rightContainer
}
