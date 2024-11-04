package widgets

import (
	"image/color"

	clr "tic-tac-toe/core/gui/colors"
	cst "tic-tac-toe/core/gui/layouts"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type LobbyEntry struct {
	widget.BaseWidget
	LobbyID      string
	Name         string
	PlayerCount  string
	OnJoinButton func()
}

func NewLobbyEntry(lobbyID, name, playerCount string, onJoinButton func()) *LobbyEntry {
	entry := &LobbyEntry{
		LobbyID:      lobbyID,
		Name:         name,
		PlayerCount:  playerCount,
		OnJoinButton: onJoinButton,
	}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (le *LobbyEntry) CreateRenderer() fyne.WidgetRenderer {
	return &LobbyEntryRenderer{entry: le}
}

type LobbyEntryRenderer struct {
	entry           *LobbyEntry
	container       *fyne.Container
	lobbyNameText   *canvas.Text
	playerCountText *canvas.Text
	joinButton      *widget.Button
	border          *canvas.Rectangle
}

func (r *LobbyEntryRenderer) Layout(size fyne.Size) {
	r.container.Resize(size)
}

func (r *LobbyEntryRenderer) MinSize() fyne.Size {
	return r.container.MinSize()
}

func (r *LobbyEntryRenderer) Refresh() {
	r.lobbyNameText.Text = r.entry.Name
	r.playerCountText.Text = r.entry.PlayerCount
	r.container.Refresh()
}

func (r *LobbyEntryRenderer) Destroy() {
	// FigureOut
}

func (r *LobbyEntryRenderer) Objects() []fyne.CanvasObject {
	return r.container.Objects
}

func NewLobbyEntryRenderer(entry *LobbyEntry) *LobbyEntryRenderer {
	lobbyNameText := canvas.NewText(entry.Name, clr.Border)
	playerCountText := canvas.NewText(entry.PlayerCount, clr.Border)
	joinButton := widget.NewButton("Join", entry.OnJoinButton)

	lobbyContainerContents := container.New(
		cst.NewRatioHLayout(0.15, 0.685, 0.02, 0.270),
		lobbyNameText,
		playerCountText,
		joinButton,
	)
	border := canvas.NewRectangle(color.Transparent)
	border.StrokeWidth = 1
	border.StrokeColor = clr.Border

	return &LobbyEntryRenderer{
		entry:           entry,
		container:       container.NewStack(border, lobbyContainerContents),
		lobbyNameText:   lobbyNameText,
		playerCountText: playerCountText,
		joinButton:      joinButton,
		border:          border,
	}
}
