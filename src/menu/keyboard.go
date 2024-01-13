package menu

import (
	"ludo/src/keyboard"

	"github.com/nsf/termbox-go"
)

func (m *menuPagesType) keyboardLoop() callback {
	kb := keyboard.KeyboardProps{EvChan: make(chan keyboard.KeyboardEvent)}
	go keyboard.ListenToKeyboard(&kb)
	for {
		e := <-kb.EvChan
		kb.Pause()
		switch e.Key {
		case termbox.KeyArrowDown:
			m.menu.handleOptNav(1)
		case termbox.KeyArrowUp:
			m.menu.handleOptNav(-1)
		case termbox.KeyArrowLeft:
			m.menu.handleSubOptNav(-1)
		case termbox.KeyArrowRight:
			m.menu.handleSubOptNav(1)
		case termbox.KeySpace:
			fallthrough
		case termbox.KeyEnter:
			if quit, callback := m.handleOptSelect(); quit {
				kb.Stop()
				return callback
			}
		case termbox.KeyEsc:
			kb.Stop()
			return nil
		}
		kb.Resume()
		m.renderMenu()
	}
}
