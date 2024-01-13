package menu

type opt struct {
	label       interface{}
	onSelect    func(*menuPagesType) (bool, callback)
	onSubOptNav func(*menu, int)
	curIdx      int
	option      []string
}

type menu struct {
	curIdx  int
	options []opt
}

func (m *menu) validateOptBounds() {
	maxOptIdx := len(m.options) - 1
	if m.curIdx < 0 {
		m.curIdx = maxOptIdx
	} else if m.curIdx == maxOptIdx+1 {
		m.curIdx = 0
	}
}

func (m *menu) handleOptNav(mag int) {
	m.curIdx += mag
	m.validateOptBounds()
}

func (m *menu) handleSubOptNav(mag int) {
	curOpt := &m.options[m.curIdx]
	if curOpt.option == nil {
		return
	}
	if curOpt.onSubOptNav != nil {
		curOpt.onSubOptNav(m, mag)
	} else {
		maxOptIdx := len(curOpt.option) - 1

		curOpt.curIdx += mag

		if curOpt.curIdx < 0 {
			curOpt.curIdx = maxOptIdx
		} else if curOpt.curIdx == maxOptIdx+1 {
			curOpt.curIdx = 0
		}
	}
}

func (m *menuPagesType) handleOptSelect() (bool, callback) {
	curMenu := m.menu
	if curMenu.options[curMenu.curIdx].onSelect != nil {
		return curMenu.options[curMenu.curIdx].onSelect(m)
	}
	return false, nil
}
