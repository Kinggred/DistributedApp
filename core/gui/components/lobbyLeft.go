package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	conf "tic-tac-toe/core/config"
	ita "tic-tac-toe/core/gameLogic/interactions"
	glo "tic-tac-toe/core/global"
	cst "tic-tac-toe/core/gui/layouts"
)

func GetLobbyLeft() *fyne.Container {

	usernameBinding := binding.NewString()
	usernameBinding.AddListener(binding.NewDataListener(func() {
		glo.PlayerInfo.Mu.Lock()
		defer glo.PlayerInfo.Mu.Unlock()
		glo.PlayerInfo.Username, _ = usernameBinding.Get()
		conf.GuiLogger.Printf("Changed player name: " + glo.PlayerInfo.Username)
	}))

	userNameWidget := widget.NewEntryWithData(usernameBinding)
	userNameWidget.SetPlaceHolder("Username")

	lobbyNameWidget := widget.NewEntry()
	lobbyNameWidget.SetPlaceHolder("Lobby Name")

	// TODO: Implement Logic
	lobbyStartButton := widget.NewButton("Host", func() {
		ita.HostAGame(lobbyNameWidget.Text)
		glo.GUIState.Mu.Lock()
		defer glo.GUIState.Mu.Unlock()
		glo.GUIState.LeftContainer.Objects = GetWaitingLeft().Objects
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
