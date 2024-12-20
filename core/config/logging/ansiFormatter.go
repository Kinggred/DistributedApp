package logging

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type AnsiFormatter struct{}

func (f *AnsiFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var (
		appColor        = "\033[32m"
		serverColor     = "\033[31m"
		clientColor     = "\033[33m"
		gameColor       = "\033[37m"
		guiColor        = "\033[36m"
		resetColor      = "\033[0m"
		additionalColor = "\033[35m"
	)

	component, ok := entry.Data[COMP]
	var componentString string

	if ok {
		switch component {
		case "App":
			componentString = fmt.Sprintf("%s[%s]%s ", appColor, component, resetColor)
		case "Server":
			componentString = fmt.Sprintf("%s[%s]%s ", serverColor, component, resetColor)
		case "Client":
			componentString = fmt.Sprintf("%s[%s]%s ", clientColor, component, resetColor)
		case "Game":
			componentString = fmt.Sprintf("%s[%s]%s ", gameColor, component, resetColor)
		case "GUI":
			componentString = fmt.Sprintf("%s[%s]%s ", guiColor, component, resetColor)
		}
		delete(entry.Data, COMP)
	}

	fields := ""
	for key, value := range entry.Data {
		fields += fmt.Sprintf(" %s%s:%s %v", additionalColor, key, resetColor, value)
	}

	timestamp := time.Now().Format("15:04:05")
	level := entry.Level.String()
	logMessage := fmt.Sprintf("%s[%s] %s: %s%s\n", componentString, level, timestamp, entry.Message, fields)
	return []byte(logMessage), nil
}
