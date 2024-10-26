package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Formatter struct{}

var ServerLogger *logrus.Entry
var ClientLogger *logrus.Entry
var AppLogger *logrus.Entry

var (
	appColor    = "\033[32m"
	serverColor = "\033[31m"
	clientColor = "\033[33m"
	resetColor  = "\033[0m"
)

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	component, ok := entry.Data["component"]
	var componentString string

	if ok {
		switch component {
		case "App":
			componentString = fmt.Sprintf("%s[%s]%s ", appColor, component, resetColor)
		case "Server":
			componentString = fmt.Sprintf("%s[%s]%s ", serverColor, component, resetColor)
		case "Client":
			componentString = fmt.Sprintf("%s[%s]%s ", clientColor, component, resetColor)
		}
		delete(entry.Data, "component")
	}

	logMessage := fmt.Sprintf("%s%s\n", componentString, entry.Message)
	return []byte(logMessage), nil
}

func SetupLoggers() {
	logger := logrus.New()
	logger.SetFormatter(&Formatter{})
	AppLogger = logger.WithField("component", "App")
	ServerLogger = logger.WithField("component", "Server")
	ClientLogger = logger.WithField("component", "Client")
}
