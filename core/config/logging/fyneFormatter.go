package logging

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type FyneFormatter struct{}

func (f *FyneFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	component, ok := entry.Data[COMP]
	var componentString string

	if ok {
		switch component {
		case "App":
			componentString = fmt.Sprintf("**[%s]** ", component)
		case "Server":
			componentString = fmt.Sprintf("**[%s]** ", component)
		case "Client":
			componentString = fmt.Sprintf("**[%s]** ", component)
		case "Game":
			componentString = fmt.Sprintf("**[%s]** ", component)
		case "GUI":
			componentString = fmt.Sprintf("**[%s]** ", component)
		}
		delete(entry.Data, COMP)
	}

	fields := ""
	for key, value := range entry.Data {
		fields += fmt.Sprintf(" *%s:* `%v`", key, value)
	}

	timestamp := time.Now().Format("15:04:05")
	level := entry.Level.String()
	logMessage := fmt.Sprintf("%s[%s] %s: %s%s\n", componentString, level, timestamp, entry.Message, fields)
	return []byte(logMessage), nil
}
