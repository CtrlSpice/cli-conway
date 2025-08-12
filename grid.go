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
func (g *Grid) SetCell(x, y int, alive bool) {
	if x >= 0 && x < g.width && y >= 0 && y < g.height {
		if alive {
			g.cells[y][x] = 1
		} else {
			g.cells[y][x] = 0
		}
	}
}

// GetCell returns the state of a cell at the specified position
func (g *Grid) GetCell(x, y int) byte {
	if x >= 0 && x < g.width && y >= 0 && y < g.height {
		return g.cells[y][x]
	}
	return 0
}

// Display renders the grid to the terminal
func (g *Grid) Display() {
	// Clear the terminal (works on most Unix-like systems)
	fmt.Print("\033[H\033[2J")

	// Print top border
	fmt.Print("┌")
	for i := 0; i < g.width*2+1; i++ {
		fmt.Print("─")
	}
	fmt.Println("┐")

	// Print grid content
	for y := 0; y < g.height; y++ {
		fmt.Print("│ ")
		for x := 0; x < g.width; x++ {
			if g.cells[y][x] == 1 {
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
	for i := 0; i < g.width*2+1; i++ {
		fmt.Print("─")
	}
	fmt.Println("┘")
}

// Randomize fills the grid with random live cells
func (g *Grid) Randomize() {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			// Simple random: use time-based seed for now
			// In a real implementation, you'd want a proper random generator
			if (x+y+int(time.Now().UnixNano()))%3 == 0 {
				g.cells[y][x] = 1
			} else {
				g.cells[y][x] = 0
			}
		}
	}
}
