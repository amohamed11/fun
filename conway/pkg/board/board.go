package board

import (
	"conway/pkg/cell"
	"fmt"
	"math/rand"
	"time"
)

type Board struct {
	cells       [][]cell.Cell
	max_x       int
	max_y       int
	current_gen int
}

func NewBoard(x int, y int) *Board {
	board := &Board{max_x: x, max_y: y, current_gen: 0}
	board.init()

	return board
}

// Tick runs a single iteration of the board
func (board *Board) Tick() {
	if board.current_gen == 0 {
		board.initialPattern()
	} else {
		board.update()
	}

	board.current_gen++
}

// Draw handles drawing the current state of the board
func (board *Board) Draw() {
	// clear the screen using escape codes
	fmt.Print("\033[H\033[2J")

	for x := 0; x < board.max_x; x++ {
		fmt.Print("|")
		for y := 0; y < board.max_y; y++ {
			fmt.Printf("%s", board.cells[x][y])
		}
		fmt.Println("|")
	}
	fmt.Printf("Generation: %d\n", board.current_gen)
}

// init initalizes the board
// fills the board with dead cells
func (board *Board) init() {
	// setup our board of cells
	board.cells = make([][]cell.Cell, board.max_x)
	for i := range board.cells {
		board.cells[i] = make([]cell.Cell, board.max_y)
	}

	// fill it with cells
	for x := 0; x < board.max_x; x++ {
		for y := 0; y < board.max_y; y++ {
			board.addCell(x, y)
		}
	}
}

// Update handles updating all the cells in the board
func (board *Board) update() {
	for x := 0; x < board.max_x; x++ {
		for y := 0; y < board.max_y; y++ {
			neighbourCount := board.getNeighbourCount(board.cells[x][y])
			board.cells[x][y].HandleUpdate(neighbourCount)
		}
	}

	for x := 0; x < board.max_x; x++ {
		for y := 0; y < board.max_y; y++ {
			if board.cells[x][y].NeedsUpdate {
				board.cells[x][y].Flip()
			}
		}
	}
}

func (board *Board) initialPattern() {
	// arbitrary selected as the number of initial cells
	count := 3

	rand.Seed(time.Now().Unix())
	for i := 0; i < count; i++ {

		// random x coord in range [3, max_x - 3]
		randX := rand.Intn((board.max_x-3)-3+1) + 3
		// random y coord in range [3, max_y - 3]
		randY := rand.Intn((board.max_y-3)-3+1) + 3

		board.cells[randX][randY].Dead = false
		board.cells[randX+1][randY+1].Dead = false
		board.cells[randX+1][randY-1].Dead = false
	}
}

func (board *Board) addCell(x int, y int) {
	newCellId := len(board.cells) + 1
	board.cells[x][y] = cell.NewCell(newCellId, x, y)
}

// getNeighbourCount takes a cell and gets the number of neighbours it has
func (board *Board) getNeighbourCount(c cell.Cell) int {
	count := 0
	neighbours := make([]*cell.Cell, 8)

	// top left
	neighbours[0] = board.getCell(c.X+1, c.Y+1)
	// top middle
	neighbours[1] = board.getCell(c.X, c.Y+1)
	// top right
	neighbours[2] = board.getCell(c.X-1, c.Y+1)
	// middle left
	neighbours[3] = board.getCell(c.X+1, c.Y)
	// middle right
	neighbours[4] = board.getCell(c.X-1, c.Y)
	// bottom left
	neighbours[5] = board.getCell(c.X+1, c.Y-1)
	// bottom middle
	neighbours[6] = board.getCell(c.X, c.Y-1)
	// bottom right
	neighbours[7] = board.getCell(c.X-1, c.Y-1)

	for _, neighbour := range neighbours {
		if neighbour != nil && !neighbour.Dead {
			count++
		}
	}

	return count
}

func (board *Board) getCell(x int, y int) *cell.Cell {
	if (x >= board.max_x || x < 0) || (y >= board.max_y || y < 0) {
		return nil
	}

	return &board.cells[x][y]
}
