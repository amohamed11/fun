package game

import (
	"conway/pkg/board"
	"time"
)

type Game struct {
	board       *board.Board
	current_gen int
}

func NewGame(x int, y int) *Game {
	board := board.NewBoard(x, y)
	game := Game{board: board, current_gen: 0}

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
