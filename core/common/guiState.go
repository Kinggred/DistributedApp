package common

import (
	"sync"

	"fyne.io/fyne/v2"
)

type MainScreenStruct struct {
	LeftContainer  *fyne.Container
	RightContainer *fyne.Container
	Mu             sync.RWMutex
}
