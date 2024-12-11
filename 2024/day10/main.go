package main

import (
	"bufio"
	"common"
	trailfinder "day10/trail_finder"
)

func main() {
	common.Solve(Part1)
	common.Solve(Part2)
}

func Part1(scanner *bufio.Scanner) int {
	trailFinder := trailfinder.NewTrailFinder(scanner)

	trails, _ := trailFinder.FindTrails()
	return trails
}

func Part2(scanner *bufio.Scanner) int {
	trailFinder := trailfinder.NewTrailFinder(scanner)

	_, trails := trailFinder.FindTrails()
	return trails
}
