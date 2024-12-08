package main

import (
	"bufio"
	"common"
	"day08/antenna_map"
	"fmt"
)

func main() {
	common.Solve(Part1)
	common.Solve(Part2)
}

func Part1(scanner *bufio.Scanner) string {
	return fmt.Sprint(antenna_map.NewAntennaMap(scanner, false).CountAntiNodes())
}

func Part2(scanner *bufio.Scanner) string {
	return fmt.Sprint(antenna_map.NewAntennaMap(scanner, true).CountAntiNodes())
}
