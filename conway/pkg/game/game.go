package game

import (
	"conway/pkg/board"
	"time"
)

type Game struct {
	board *board.Board
}

func NewGame(x int, y int) *Game {
	board := board.NewBoard(x, y)
	game := Game{board: board}

	return &game
}

func (game *Game) Run() {
	for {
		// let board handle tick
		game.board.Tick()

		// re-draw the frame
		game.board.Draw()

		// update window at 1 fps
		time.Sleep(1 * time.Second)
	}
}
