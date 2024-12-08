package main

import (
	"bufio"
	"common"
	"day08/antenna_map"
)

func main() {
	common.Solve(Part1)
	common.Solve(Part2)
}

func Part1(scanner *bufio.Scanner) int {
	return antenna_map.NewAntennaMap(scanner, false).CountAntiNodes()
}

func Part2(scanner *bufio.Scanner) int {
	return antenna_map.NewAntennaMap(scanner, true).CountAntiNodes()
}
