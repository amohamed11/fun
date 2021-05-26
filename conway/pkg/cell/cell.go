package cell

import "fmt"

type Cell struct {
	Id   int
	Dead bool
	X    int
	Y    int
}

func NewCell(id int, x int, y int) Cell {
	cell := Cell{Id: id, X: x, Y: y, Dead: true}

	return cell
}

// HandleUpdate checks conditions of life for a cell
// Rules (from wikipedia):
//  1. Any live cell with two or three live neighbours survives.
//  2. Any dead cell with three live neighbours becomes a live cell.
//  3. All other live cells die in the next generation. Similarly, all other dead cells stay dead.
func (cell *Cell) HandleUpdate(neighbourCount int) {
	if neighbourCount == 2 || neighbourCount == 3 {
		cell.Dead = false
	} else {
		cell.Dead = true
	}
}

func (cell Cell) String() string {
	if cell.Dead {
		return fmt.Sprint("•")
	} else {
		return fmt.Sprint("■")
	}
}
