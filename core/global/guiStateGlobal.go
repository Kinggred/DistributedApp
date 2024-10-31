package gloabals

import (
	com "tic-tac-toe/core/common"

	"fyne.io/fyne/v2/container"
)

var GUIState com.MainScreenStruct

func InitGUIStateGlobal() {
	GUIState = com.MainScreenStruct{
		LeftContainer:  container.NewVBox(),
		RightContainer: container.NewVBox(),
	}
}
