package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	width  int
	height int
	cells  string
	random bool
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "cli-conway",
		Short: "A terminal-based Conway's Game of Life implementation",
		Long: `A CLI application written in Go that implements Conway's Game of Life on the terminal. 
		It serves no purpose, but it sure it fun to play with.`,
		Run: run,
	}

	// Add flags
	rootCmd.Flags().IntVarP(&width, "width", "x", 42, "Grid width")
	rootCmd.Flags().IntVarP(&height, "height", "y", 42, "Grid height")
	rootCmd.Flags().StringVarP(&cells, "cells", "c", "[[1,0],[2,1],[0,2],[1,2],[2,2]]", "Start with live cells as JSON array: '[[x1,y1],[x2,y2],...]'")
	rootCmd.Flags().BoolVarP(&random, "random", "r", false, "Randomize your start state")

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
	}
}

func run(cmd *cobra.Command, args []string) {
	// Create a grid with the specified dimensions
	grid := NewGrid(width, height)

	if random {
		// Use random initial state
		grid.Randomize()
	} else {
		// Parse and set initial cells from JSON
		var cellCoords [][]int
		if err := json.Unmarshal([]byte(cells), &cellCoords); err != nil {
			fmt.Printf("Error parsing cells: %v\n", err)
			return
		}

		for _, coord := range cellCoords {
			if len(coord) == 2 {
				x, y := coord[0], coord[1]
				// Check if coordinates are within grid bounds
				if x < 0 || x >= width || y < 0 || y >= height {
					fmt.Printf("Warning: Cell coordinate [%d,%d] is outside grid bounds (%dx%d). Unceremoniously skipping it.\n", x, y, width, height)
					continue
				}
				grid.SetCell(x, y, true)
			}
		}
	}

	// Display the grid
	grid.Display()

	fmt.Printf("Grid %dx%d displayed! Press Enter to exit...\n", width, height)
	fmt.Scanln() // Wait for user input
}
