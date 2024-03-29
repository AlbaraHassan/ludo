package arena

import (
	board "ludo/src/ludo-board"
	tbu "ludo/src/termbox-utils"
	"time"

	"github.com/nsf/termbox-go"
)

func (a *Arena) CurPawn() board.Pawn {
	return a.CurPlayer().Pawns[a.Board.CurPawnIdx]
}

func (a *Arena) CurPlayer() board.Player {
	return a.Board.Players[a.CurTurn]
}

func (a *Arena) MakeMove(sleepTime time.Duration, doRender bool) (hasDestroyed bool, hasReachedDest bool) {
	by := a.Dice.Value
	curPlayerColor := a.CurPlayer().Color
	curPawn := a.CurPawn()
	if curPawn["homeNode"] == curPawn["curNode"] && by != 6 {
		return
	}
	if curPawn["homeNode"] == curPawn["curNode"] {
		by = 1
	}
	for i := 0; i < by; i++ {
		curPawn["curNode"].Cell.Bg = termbox.ColorDefault

		if toDest := curPawn["curNode"].Next["toDest"]; toDest != nil && toDest.Cell.Fg == curPlayerColor {
			curPawn.MoveToNext("toDest", curPlayerColor)
		} else if curPawn["curNode"].Next["common"] != nil {
			curPawn.MoveToNext("common", curPlayerColor)
		} else {
			break
		}

		if doRender {
			a.Render()
		}

		time.Sleep(sleepTime)
	}

	hasDestroyed = a.checkDestroy()
	hasReachedDest = curPawn.IsAtDest()

	if doRender {
		a.Render()
	}

	return
}

func (a *Arena) checkDestroy() (hasDestroyed bool) {
	curCell := a.CurPawn()["curNode"].Cell

	for i, p := range a.Board.Players {
		if i == a.CurTurn || !p.IsParticipant() {
			continue
		}

		for j, _pawn := range p.Pawns {
			c := _pawn["curNode"].Cell
			if c.X == curCell.X && c.Y == curCell.Y {
				hasDestroyed = true
				if a.CurPlayer().IsBot() {
					a.ResetBotPawn(i, j)
				}
				_pawn["curNode"] = _pawn["homeNode"]
			}
		}
	}

	return
}

func (a *Arena) IsGameOver() bool {
	return a.NextWinningPos >= a.ParticipantsCount-1
}

func (a *Arena) RepaintCurPawn() {
	curCell := a.CurPawn()["curNode"].Cell
	tbu.SetBg(curCell.X, curCell.Y, a.Board.Players[a.CurTurn].Color)
}
