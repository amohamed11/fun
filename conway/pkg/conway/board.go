package conway

type Board struct {
	cells       [][]Cell
	current_gen int
	max_x       int
	max_y       int
}

func NewBoard(x int, y int) *Board {
	board := Board{max_x: x, max_y: y, current_gen: 0}
	board.init()

	return board
}

// Tick runs a single iteration of the board
func (board *Board) Tick() {
	for _, c := range board.cells {
		neighbourCount := board.getNeighbourCount(c)
		c.HandleUpdate(neighbourCount)
	}

	current_gen += 1
}

// init initalizes the board
// fills the board with dead cells, then randomly selects 3 starting cells
func (board *Board) init() {
	// setup our board of cells
	board.cells = make([][]Cell, board.max_x)
	for i := range board.cells {
		board.cells[i] = make([]Cell, board.max_y)
	}

	// fill it with cells
	for x := 0; x < board.max_x; x++ {
		for y := 0; y < board.max_y; y++ {
			board.addCell(x, y)
		}
	}
}

func (board *Board) addCell(x int, y int) {
	newCellId := len(board.cells) + 1
	board.cells[x][y] = Cell.NewCell(newCellId, x, y)
}

// getNeighbourCount takes a cell and gets the number of neighbours it has
func (board *Board) getNeighbourCount(cell *Cell) int {
	count := 0
	neighbours := make([]Cell, 8)

	// top left
	neighbours[0] = board.getCell(x+1, y+1)
	// top middle
	neighbours[1] = board.getCell(x, y+1)
	// top right
	neighbours[2] = board.getCell(x-1, y+1)
	// middle left
	neighbours[3] = board.getCell(x+1, y)
	// middle right
	neighbours[4] = board.getCell(x-1, y)
	// bottom left
	neighbours[5] = board.getCell(x+1, y-1)
	// bottom middle
	neighbours[6] = board.getCell(x, y-1)
	// bottom right
	neighbours[7] = board.getCell(x-1, y-1)

	for _, neighbour := range neighbours {
		if neighbour != nil && !neighbour.dead {
			count++
		}
	}

	return count
}

func (board *Board) getCell(x int, y int) *Cell {
	if (x >= board.max_x || x < 0) || (y >= board.max_y || y < 0) {
		return nil
	}

	return board.cells[x][y]
}
