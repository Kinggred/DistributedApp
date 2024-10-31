package logging

import (
	glo "tic-tac-toe/core/global"

	"github.com/sirupsen/logrus"
)

type LogHook struct{}

func (h *LogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *LogHook) Fire(entry *logrus.Entry) error {
	entryCopy := &logrus.Entry{
		Logger:  entry.Logger,
		Message: entry.Message,
		Data:    make(logrus.Fields),
		Level:   entry.Level,
		Time:    entry.Time,
	}

	for key, value := range entry.Data {
		entryCopy.Data[key] = value
	}

	formatter := &FyneFormatter{}
	msg, _ := formatter.Format(entryCopy)
	glo.LogOutput.Set(string(msg))
	return nil
}
