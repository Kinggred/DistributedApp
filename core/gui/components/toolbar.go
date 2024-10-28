package components

import (
	clr "tic-tac-toe/core/gui/colors"
	cst "tic-tac-toe/core/gui/layouts"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func GetToolbar() fyne.CanvasObject {
	versionText := canvas.NewText("App Version: ", clr.Green)
	toolbarContainer := container.New(cst.NewRatioHLayout(0.3, 0.7), versionText, spacer)
	return toolbarContainer
}
