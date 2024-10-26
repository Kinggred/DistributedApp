package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var green color.NRGBA = color.NRGBA{R: 0, G: 180, B: 0, A: 255}

func RenderStartingScreen() {
	app := app.New()
	window := app.NewWindow("Starting Screen")

	userNameWidget := widget.NewEntry()
	userNameWidget.SetPlaceHolder("Username")
	spacer := layout.NewSpacer()
	lobbyNameWidget := widget.NewEntry()
	lobbyNameWidget.SetPlaceHolder("Lobby Name")
	lobbyStartButton := widget.NewButton("Host", func() {})
	lobbyCreationContainer := container.New(NewRatioHLayout(0.685, 0.01, 0.285), lobbyNameWidget, spacer, lobbyStartButton)
	leftContainer := container.New(layout.NewVBoxLayout(), userNameWidget, spacer, lobbyCreationContainer)

	refreshButton := widget.NewButton("Refresh", func() {})
	joinButton := widget.NewButton("Join", func() {})

	lobbyNameText := canvas.NewText("Lobby Name", green)
	playerCountText := canvas.NewText("1/2", green)
	lobbyContainerContents := container.New(NewRatioHLayout(0.15, 0.685, 0.02, 0.270), spacer, lobbyNameText, spacer, playerCountText)
	lobbyContainerContentsPadded := container.NewPadded(lobbyContainerContents)
	border := canvas.NewRectangle(color.Transparent)
	border.StrokeWidth = 1
	border.StrokeColor = green
	hiddenButton := widget.NewButton("", func() {})

	//hoverEffect := func(isHovered bool) {
	//	if isHovered {
	//		border.StrokeColor = color.Transparent
	//	} else {
	//		border.StrokeColor = green
	//	}
	//	border.Refresh()
	//}

	lobbyContainer := container.NewStack(hiddenButton, border, lobbyContainerContentsPadded)

	buttonRow := container.New(NewRatioHLayout(0.685, 0.01, 0.285), joinButton, spacer, refreshButton)
	onlyOpenCheckbox := widget.NewCheck("Only open", func(b bool) {})
	onlyLocalCheckbox := widget.NewCheck("Only local", func(b bool) {})
	checkboxRow := container.New(NewRatioHLayout(0.45, 0.1, 0.45), onlyOpenCheckbox, spacer, onlyLocalCheckbox)
	rightContainer := container.New(layout.NewVBoxLayout(), lobbyContainer, spacer, buttonRow, checkboxRow)

	versionText := canvas.NewText("App Version: ", green)
	toolBarContainer := container.New(NewRatioHLayout(0.3, 0.7), versionText, spacer)

	mainContainer := container.New(NewRatioHLayout(0.45, 0.1, 0.45), leftContainer, spacer, rightContainer)
	combinedContainer := container.New(NewRatioVLayout(0.944, 0.005, 0.06), mainContainer, spacer, toolBarContainer)
	paddedContainer := container.NewPadded(combinedContainer)

	window.SetContent(paddedContainer)
	window.Resize(fyne.NewSize(550, 300))
	window.ShowAndRun()
}
