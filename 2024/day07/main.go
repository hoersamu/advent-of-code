package main

import (
	"bufio"
	"common"
	"day07/solver"
)

func main() {
	common.Solve(Part1)
	common.Solve(Part2)
}

func Part1(scanner *bufio.Scanner) int {
	return solver.NewEquationSolver().SolveEquations(scanner, false)
}

func Part2(scanner *bufio.Scanner) int {
	return solver.NewEquationSolver().SolveEquations(scanner, true)
}
