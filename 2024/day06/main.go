package main

import (
	"bufio"
	"common"
	"slices"
)

func main() {
	common.Solve(Part1)
	common.Solve(Part2)
}

const CHARACTER = "^"
const WALL = "#"
const VISITED = "X"

var matrix = [][]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func findInitialPosition(grid [][]string) (int, int) {
	for y, row := range grid {
		x := slices.Index(row, CHARACTER)
		if x != -1 {
			return x, y
		}
	}
	return 0, 0
}

func getNextPosition(xIndex, yIndex int, direction int) (int, int) {
	return xIndex + matrix[direction][0], yIndex + matrix[direction][1]
}

type GridRunner struct {
	grid      [][]string
	visited   [][]bool
	gridSize  common.Grid
	xIndex    int
	yIndex    int
	direction int
}

func newGridRunner(grid [][]string) *GridRunner {
	xSize, ySize := len(grid[0]), len(grid)
	xIndex, yIndex := findInitialPosition(grid)

	visited := make([][]bool, ySize)
	for i := range visited {
		visited[i] = make([]bool, xSize)
	}

	return &GridRunner{
		grid:      grid,
		visited:   visited,
		gridSize:  common.Grid{SizeX: xSize, SizeY: ySize},
		xIndex:    xIndex,
		yIndex:    yIndex,
		direction: 0,
	}
}

func (r *GridRunner) run(checkLoop bool) (int, bool) {
	visitedFields := 0
	if checkLoop {
		// Reset visited array for loop checking
		for y := range r.visited {
			for x := range r.visited[y] {
				r.visited[y][x] = false
			}
		}
	}

	type state struct {
		x, y, dir int
	}
	positions := make(map[state]bool)

	for {
		// Mark current position
		if !checkLoop {
			if r.grid[r.yIndex][r.xIndex] != VISITED {
				visitedFields++
				r.grid[r.yIndex][r.xIndex] = VISITED
			}
		} else {
			currentState := state{r.xIndex, r.yIndex, r.direction}
			if positions[currentState] {
				return 0, true
			}
			positions[currentState] = true
			r.visited[r.yIndex][r.xIndex] = true
		}

		// Calculate next position
		nextX, nextY := getNextPosition(r.xIndex, r.yIndex, r.direction)

		// Check for out of bounds
		if !r.gridSize.Validate(common.Point{X: nextX, Y: nextY}) {
			return visitedFields, false
		}

		// Check for collision
		if r.grid[nextY][nextX] == WALL {
			// Turn right
			r.direction = (r.direction + 1) % 4
			continue
		}

		r.xIndex, r.yIndex = nextX, nextY
	}
}

func runOverGrid(grid [][]string) int {
	runner := newGridRunner(grid)
	fields, _ := runner.run(false)
	return fields
}

func isLoop(grid [][]string, visited [][]bool) bool {
	runner := newGridRunner(grid)
	runner.visited = visited // Reuse existing visited array
	_, isLoop := runner.run(true)
	return isLoop
}

func Part1(scanner *bufio.Scanner) int {
	grid := common.ScanWithDelimiters(scanner, "")

	return runOverGrid(grid)
}

func Part2(scanner *bufio.Scanner) int {
	originalGrid := common.ScanWithDelimiters(scanner, "")
	xSize, ySize := len(originalGrid[0]), len(originalGrid)

	// Create visited array once and reuse
	visited := make([][]bool, ySize)
	for i := range visited {
		visited[i] = make([]bool, xSize)
	}

	// Create working grid
	grid := make([][]string, ySize)
	for i := range grid {
		grid[i] = make([]string, xSize)
	}

	// First pass to mark initial path
	// Copy original grid to working grid
	for i := range originalGrid {
		copy(grid[i], originalGrid[i])
	}

	startX, startY := findInitialPosition(grid)
	runOverGrid(grid)

	// Store the path grid state
	pathGrid := make([][]string, ySize)
	for i := range grid {
		pathGrid[i] = make([]string, xSize)
		copy(pathGrid[i], grid[i])
	}

	loops := 0
	for y, row := range pathGrid {
		for x, field := range row {
			if field == VISITED && !(x == startX && y == startY) {
				// Reset grid to original state
				for i := range originalGrid {
					copy(grid[i], originalGrid[i])
				}
				// Add test block
				grid[y][x] = WALL

				if isLoop(grid, visited) {
					loops++
				}
			}
		}
	}

	return loops
}
