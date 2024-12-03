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

func isLineSafe(line []int, isLineAscending bool, skipIndex int) bool {
	isSafe := true

	lineWithoutIndex := line
	if skipIndex != -1 {
		lineWithoutIndex = common.RemoveIndex(line, skipIndex)
	}

	for i := 0; i < len(lineWithoutIndex)-1; i++ {
		multiplier := 1
		if !isLineAscending {
			multiplier = -1
		}
		num1 := lineWithoutIndex[i] * multiplier
		num2 := lineWithoutIndex[i+1] * multiplier

		distance := common.GetDistance(num1, num2)

		if num1 > num2 || distance < 1 || distance > 3 {
			isSafe = false
			break
		}
	}
	return isSafe
}

func Part1(scanner *bufio.Scanner) string {
	records := common.ScanWithDelimitersAsInt(scanner, " ")

	safeCount := 0

	for _, line := range records {
		isLineAscending := line[0] < line[1]

		isSafe := isLineSafe(line, isLineAscending, -1)
		if isSafe {
			safeCount++
		}
	}

	return strconv.Itoa(safeCount)
}

func Part2(scanner *bufio.Scanner) string {
	records := common.ScanWithDelimitersAsInt(scanner, " ")

	safeCount := 0

	for _, line := range records {
		for i := -1; i < len(line); i++ {
			isSafe := isLineSafe(line, true, i) || isLineSafe(line, false, i)
			if isSafe {
				safeCount++
				break
			}
		}
	}

	return strconv.Itoa(safeCount)
}
