package main

import (
	"bufio"
	"common"
	"strconv"
)

func main() {
	common.Solve(Part1)
	common.Solve(Part2)
}

func Part1(scanner *bufio.Scanner) string {
	grid := common.ScanToLines(scanner)
	count := 0
	rows := len(grid)
	cols := len(grid[0])
	word := "XMAS"
	wordLen := len(word)

	// Single loop to check all directions
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			directions := [][2]int{
				{0, 1},  // right
				{1, 0},  // down
				{1, 1},  // down-right
				{-1, 1}, // up-right
			}

			for _, dir := range directions {
				if row+(wordLen-1)*dir[0] >= 0 && row+(wordLen-1)*dir[0] < rows &&
					col+(wordLen-1)*dir[1] >= 0 && col+(wordLen-1)*dir[1] < cols {
					var str string
					for i := 0; i < wordLen; i++ {
						str += string(grid[row+i*dir[0]][col+i*dir[1]])
					}
					if str == word || common.ReverseString(str) == word {
						count++
					}
				}
			}
		}
	}
	return strconv.Itoa(count)
}

func Part2(scanner *bufio.Scanner) string {
	grid := common.ScanToLines(scanner)
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	// Helper function to check if two positions form an M-S pair
	isValidPair := func(pos1, pos2 byte) bool {
		return (pos1 == 'S' && pos2 == 'M') || (pos1 == 'M' && pos2 == 'S')
	}

	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			if grid[row][col] != 'A' {
				continue
			}

			// Check both diagonals
			diagonal1 := isValidPair(grid[row-1][col-1], grid[row+1][col+1])
			diagonal2 := isValidPair(grid[row-1][col+1], grid[row+1][col-1])

			if diagonal1 && diagonal2 {
				count++
			}
		}
	}
	return strconv.Itoa(count)
}
