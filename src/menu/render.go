package menu

import (
	tbu "ludo/src/termbox-utils"

	"github.com/nsf/termbox-go"
)

var (
	paddingMag = 3
)

func renderSubOpts(op opt, x, y, oidx, menuIdx int) {
	var fg, bg termbox.Attribute

	if menuIdx == oidx {
		bg, fg = termbox.ColorYellow, termbox.ColorBlack
	}

	switch l := op.label.(type) {
	case string:
		x += len(l)
		tbu.RenderText(tbu.Text{Text: l, X: x, Y: y, Fg: fg, Bg: bg, InlinePadding: paddingMag})
	case termbox.Attribute:
		tbu.RenderText(tbu.Text{Text: "", X: x, Y: y, Fg: fg, Bg: l, InlinePadding: 1})
	}
	if menuIdx == oidx {
		fg, bg = termbox.ColorLightBlue, termbox.ColorDefault
	} else {
		fg, bg = termbox.ColorGreen, termbox.ColorDefault
	}
	tbu.RenderText(tbu.Text{Text: op.option[op.curIdx], X: x + 4, Y: y, Fg: fg, Bg: bg})
}

func (m *menuPagesType) renderMenu() {
	tbu.Clear()

	curMenu := m.menus[m.curIdx]

	y := m.displayPos.y

	var fg, bg termbox.Attribute

	for oidx, op := range curMenu.options {
		fg, bg = termbox.ColorRed, termbox.ColorDefault

		if op.option != nil {
			renderSubOpts(op, m.displayPos.x, y, oidx, curMenu.curIdx)
		} else {
			if curMenu.curIdx == oidx {
				bg, fg = termbox.ColorLightGray, termbox.ColorBlack
			}
			switch l := op.label.(type) {
			case string:
				tbu.RenderText(tbu.Text{Text: l, X: m.displayPos.x, Y: y, Fg: fg, Bg: bg, InlinePadding: paddingMag})
			case termbox.Attribute:
				tbu.RenderText(tbu.Text{Text: "", X: m.displayPos.x, Y: y, Fg: fg, Bg: l, InlinePadding: 1})
			}
		}
		y += 2
	}

	termbox.Flush()
}
