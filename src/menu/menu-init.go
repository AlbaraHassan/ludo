package menu

import "github.com/nsf/termbox-go"

type pos struct {
	x, y int
}

type menuPagesType struct {
	curIdx     int
	menus      []*menu
	displayPos pos
}

func (m *menuPagesType) changeMenuPage(pageIdx int) {
	m.curIdx = pageIdx
	m.menus[m.curIdx].curIdx = 0
}

func InitMenu() {
	menuPages := menuPagesType{
		displayPos: pos{60, 5},
		curIdx:     0,
		menus:      []*menu{&mainMenu},
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
