package main

import (
	"fmt"
	"time"
)

// Grid represents the game board
type Grid struct {
	width  int
	height int
	cells  [][]byte
}

// NewGrid creates a new grid with the specified dimensions
func NewGrid(width, height int) *Grid {
	cells := make([][]byte, height)
	for i := range cells {
		cells[i] = make([]byte, width)
	}
	return &Grid{
		width:  width,
		height: height,
		cells:  cells,
	}
}

// SetCell sets a cell at the specified position
func (grid *Grid) SetCell(x, y int, value byte) {
	if x >= 0 && x < grid.width && y >= 0 && y < grid.height {
		grid.cells[y][x] = value
	}
}

// GetCell returns the state of a cell at the specified position
func (grid *Grid) GetCell(x, y int) byte {
	if x >= 0 && x < grid.width && y >= 0 && y < grid.height {
		return grid.cells[y][x]
	}
	return 0
}

// MakeItSo renders the grid to the terminal
func (grid *Grid) MakeItSo() {
	// Move cursor to top-left without clearing screen
	fmt.Print("\033[H")

	// Print top border
	fmt.Print("┌")
	for i := 0; i < grid.width*2+1; i++ {
		fmt.Print("─")
	}
	fmt.Println("┐")

	// Print grid content
	for y := 0; y < grid.height; y++ {
		fmt.Print("│ ")
		for x := 0; x < grid.width; x++ {
			if grid.cells[y][x] == 1 {
				fmt.Print("█") // Live cell
			} else {
				fmt.Print(" ") // Dead cell
			}
			fmt.Print(" ")
		}
		fmt.Println("│")
	}

	// Print bottom border
	fmt.Print("└")
	for i := 0; i < grid.width*2+1; i++ {
		fmt.Print("─")
	}
	fmt.Println("┘")
}

// Randomize fills the grid with random live cells
func (grid *Grid) Randomize() {
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			// Simple random: use time-based seed. It's good enough.
			if (x+y+int(time.Now().UnixNano()))%3 == 0 {
				grid.cells[y][x] = 1
			} else {
				grid.cells[y][x] = 0
			}
		}
	}
}

// Boldly generates "The Next Generation" (Get it? Get it? I will show myself out) of grid
func (grid *Grid) BoldlyGo() *Grid {
	// Create a new grid for the next generation
	nextGen := NewGrid(grid.width, grid.height)
	
	// Apply Conway's rules to each cell
	for y := range grid.cells {
		for x := range grid.cells[y] {
			lifeformCount := grid.scanForLifeforms(x, y)
			// If the cell is alive
			if grid.cells[y][x] == 1 {
				// Kill it if it's lonely or overcrowded
				if lifeformCount < 2 || lifeformCount > 3 {
					nextGen.SetCell(x, y, 0)
				} else {
					nextGen.SetCell(x, y, 1)
				}
			// If the cell is dead
			} else {
				// Reproduce if there are exactly three lifeforms in the neighborhood
				if lifeformCount == 3 {
					nextGen.SetCell(x, y, 1)
				} else {
					nextGen.SetCell(x, y, 0)
				}
			}
		}
	}
	
	return nextGen
}

// scanForLifeforms counts the number of live neighbors for a given cell
// Data loves scanning for lifeforms
func (grid *Grid) scanForLifeforms(x, y int) int {
	// Data loves scanning for lifeforms
	lifeformCount := 0
	

// Neighbor positions around cell [x][y]:
// [x-1][y+1] [x][y+1] [x+1][y+1]  (top row)
// [x-1][y]   [x][y]   [x+1][y]    (middle row - center is the cell itself)
// [x-1][y-1] [x][y-1] [x+1][y-1]  (bottom row)
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			// Skip the center cell itself
			if dx == 0 && dy == 0 {
				continue
			}
			
			newX := x + dx
			newY := y + dy
			if newX >= 0 && newX < grid.width && newY >= 0 && newY < grid.height {
				lifeformCount += int(grid.cells[newY][newX])
			}
		}
	}

	return lifeformCount
}

