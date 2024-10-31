package logging

import (
	glo "tic-tac-toe/core/global"

	"fyne.io/fyne/v2/data/binding"
	"github.com/sirupsen/logrus"
)

var ServerLogger *logrus.Entry
var ClientLogger *logrus.Entry
var AppLogger *logrus.Entry
var GameLogger *logrus.Entry
var GuiLogger *logrus.Entry

const COMP string = "component"

func SetupLoggers() {
	glo.LogOutput = binding.NewString()
	logger := logrus.New()
	logger.AddHook(&LogHook{})
	logger.SetFormatter(&AnsiFormatter{})
	AppLogger = logger.WithField(COMP, "App")
	ServerLogger = logger.WithField(COMP, "Server")
	ClientLogger = logger.WithField(COMP, "Client")
	GameLogger = logger.WithField(COMP, "Game")
	GuiLogger = logger.WithField(COMP, "GUI")
}
