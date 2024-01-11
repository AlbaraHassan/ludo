package menu

import (
	"github.com/nsf/termbox-go"
	"ludo/src/common"
	"ludo/src/game"
)

type callback func() int

var (
	options = []string{"Player", "Bot", "-"}

	mainMenu = menu{
		options: []opt{
			{
				label:  termbox.ColorYellow,
				option: options,
			},
			{
				label:  termbox.ColorBlue,
				option: options,
			},
			{
				label:  termbox.ColorRed,
				option: options,
			},
			{
				label:  termbox.ColorGreen,
				option: options,
			},
			{
				label: "Done",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					players := []common.PlayerData{}
					curMenuOpts := mpt.menus[mpt.curIdx].options
					curMenuOpts = curMenuOpts[:len(curMenuOpts)-2]
					for _, opt := range curMenuOpts {
						players = append(players, common.PlayerData{Color: opt.label.(termbox.Attribute), Type: opt.option[opt.curIdx]})
					}
					return true, func() int {
						termbox.Close()
						err := termbox.Init()
						if err != nil {
							return 0
						}
						game.StartGame(players)
						return 0
					}
				},
			},
			{
				label: "Exit",
				onSelect: func(mpt *menuPagesType) (bool, callback) {
					return true, nil
				},
			},
		},
	}
)
