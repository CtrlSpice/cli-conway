# CLI-Conway
A terminal-based implementation of Conway's Game of Life in Go. 
I'm just messing around with it as part of my time at the Recurse Center.

## Project Structure
- `main.go` - Contains the basic grid infrastructure
- `go.mod` - Go module definition

## Conway's Rules

1. Any live cell with fewer than 2 live neighbors dies (underpopulation)
2. Any live cell with 2 or 3 live neighbors lives on to the next generation
3. Any live cell with more than 3 live neighbors dies (overpopulation)
4. Any dead cell with exactly 3 live neighbors becomes a live cell (reproduction)
