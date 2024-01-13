package menu

import "github.com/nsf/termbox-go"

type pos struct {
	x, y int
}

type menuPagesType struct {
	curIdx     int
	menu       *menu
	displayPos pos
}

func InitMenu() {
	menuPages := menuPagesType{
		displayPos: pos{60, 5},
		menu:       &mainMenu,
	}
	menuPages.renderMenu()
	if callback := menuPages.keyboardLoop(); callback != nil {
		exitCode := callback()
		switch exitCode {
		case 0:
			return
		case -1:
			if !termbox.IsInit {
				termbox.Init()
			}
			InitMenu()
		}
	}
}
