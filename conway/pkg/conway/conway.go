package conway

import "time"

type Game struct {
	board       *Board
	current_gen int
}

func NewGame(x int, y int) *Game {
	board := NewBoard(x, y)
	game := Game{board: board, current_gen: 0}

	return game
}

func (game *Game) Run() {
	for {
		// let board handle tick
		game.tick()

		// re-draw the frame
		game.draw()

		// update window at 15 fps
		time.Sleep(1 / 15 * time.Second)
	}
}

// tick runs a single iteration of the board
func (game *Game) tick() {
	if current_gen == 0 {
		game.board.randInitialPattern()
	} else {
		for _, c := range game.board.cells {
			neighbourCount := game.board.getNeighbourCount(c)
			c.HandleUpdate(neighbourCount)
		}
	}

	current_gen += 1
}

func (game *Game) draw() {
	for i := 0; i < len(game.board); i++ {
		fmt.printf("|%b|", game.board.cells[i])
	}
}
