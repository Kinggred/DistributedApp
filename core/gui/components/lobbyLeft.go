package components

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	log "tic-tac-toe/core/config/logging"
	ita "tic-tac-toe/core/gameLogic/interactions"
	glo "tic-tac-toe/core/global"
	cst "tic-tac-toe/core/gui/layouts"
)

func GetLobbyLeft(ctx context.Context) *fyne.Container {

	usernameBinding := binding.NewString()
	usernameBinding.AddListener(binding.NewDataListener(func() {
		username, _ := usernameBinding.Get()
		if glo.LocalPlayer.GetUsername() != username {
			glo.LocalPlayer.SetUsername(username)
			log.GuiLogger.Printf("Changed player name: " + glo.LocalPlayer.GetUsername())
		}
	}))

	userNameWidget := widget.NewEntryWithData(usernameBinding)
	userNameWidget.SetPlaceHolder("Username")
	if glo.LocalPlayer.GetUsername() != "" {
		userNameWidget.SetText(glo.LocalPlayer.GetUsername())
	}

	lobbyNameWidget := widget.NewEntry()
	lobbyNameWidget.SetPlaceHolder("Lobby Name")

	lobbyStartButton := widget.NewButton("Host", func() {
		ita.HostAGame(lobbyNameWidget.Text)
		glo.GUIState.Mu.Lock()
		defer glo.GUIState.Mu.Unlock()
		glo.GUIState.LeftContainer.Objects = GetWaitingLeft(ctx).Objects
		glo.GUIState.RightContainer.Objects = GetWaitingRight().Objects
	})

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
