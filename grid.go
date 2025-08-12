package main

import (
	"fmt"
	"time"
)

// Grid represents the game board using a flattened bitmask approach
type Grid struct {
	width  int
	height int
	cells  []uint64 // Flattened grid where each uint64 represents 64 cells
}

// NewGrid creates a new grid with the specified dimensions
func NewGrid(width, height int) *Grid {
	// Calculate how many uint64s we need to store all cells
	// Each uint64 can store 64 cells, so we need (width * height + 63) / 64
	totalCells := width * height
	numUint64s := (totalCells + 63) / 64
	
	cells := make([]uint64, numUint64s)
	return &Grid{
		width:  width,
		height: height,
		cells:  cells,
	}
}

// getBitIndex converts x,y coordinates to bit position in the flattened array
func (grid *Grid) getBitIndex(x, y int) (uint64Index int, bitPos uint) {
	linearIndex := y*grid.width + x
	uint64Index = linearIndex / 64
	bitPos = uint(linearIndex % 64)
	return
}

// SetCell sets a cell at the specified position using bitwise operations
func (grid *Grid) SetCell(x, y int, value byte) {
	if x >= 0 && x < grid.width && y >= 0 && y < grid.height {
		uint64Index, bitPos := grid.getBitIndex(x, y)
		if value == 1 {
			grid.cells[uint64Index] |= (1 << bitPos) // Set bit
		} else {
			grid.cells[uint64Index] &^= (1 << bitPos) // Clear bit
		}
	}
}

// GetCell returns the state of a cell at the specified position
func (grid *Grid) GetCell(x, y int) byte {
	if x >= 0 && x < grid.width && y >= 0 && y < grid.height {
		uint64Index, bitPos := grid.getBitIndex(x, y)
		if (grid.cells[uint64Index] & (1 << bitPos)) != 0 {
			return 1
		}
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
			if grid.GetCell(x, y) == 1 {
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
				grid.SetCell(x, y, 1)
			} else {
				grid.SetCell(x, y, 0)
			}
		}
	}
}

// Boldly generates "The Next Generation" using bitwise operations
func (grid *Grid) BoldlyGo() *Grid {
	// Create a new grid for the next generation
	nextGen := NewGrid(grid.width, grid.height)
	
	// Apply Conway's rules to each cell
	for y := 0; y < grid.height; y++ {
		for x := 0; x < grid.width; x++ {
			lifeformCount := grid.scanForLifeforms(x, y)
			currentCell := grid.GetCell(x, y)
			
			// If the cell is alive
			if currentCell == 1 {
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

// scanForLifeforms counts the number of live neighbors using bitwise operations
// Data loves scanning for lifeforms
func (grid *Grid) scanForLifeforms(x, y int) int {
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
				lifeformCount += int(grid.GetCell(newX, newY))
			}
		}
	}

	return lifeformCount
}

