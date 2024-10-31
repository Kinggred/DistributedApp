package components

import (
	glo "tic-tac-toe/core/global"
	clr "tic-tac-toe/core/gui/colors"
	cst "tic-tac-toe/core/gui/layouts"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func GetToolbar(app fyne.App) fyne.CanvasObject {
	// One day I will remember this exists
	versionText := canvas.NewText("App Version: ", clr.Border)

	logLabel := widget.NewRichTextFromMarkdown("")
	glo.LogOutput.AddListener(binding.NewDataListener(func() {
		currentText, _ := glo.LogOutput.Get()
		logLabel.AppendMarkdown(currentText)

	}))

	// Logs dont support ANSI color encoding. Maybe one day I will have too much time
	logButton := widget.NewButton("logs", func() {
		logWindow := app.NewWindow("Logs")
		scrollableContainer := container.NewScroll(logLabel)
		logWindow.SetContent(scrollableContainer)
		logWindow.Resize(fyne.NewSize(400, 200))
		logWindow.Show()
	})

	toolbarContainer := container.New(cst.NewRatioHLayout(0.3, 0.6, 0.1), versionText, spacer, logButton)
	return toolbarContainer
}
