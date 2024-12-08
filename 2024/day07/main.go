package main

import (
	"bufio"
	"common"
	"day07/solver"
	"strconv"
)

func main() {
	common.Solve(Part1)
	common.Solve(Part2)
}

func Part1(scanner *bufio.Scanner) string {
	return strconv.Itoa(solver.NewEquationSolver().SolveEquations(scanner, false))
}

func Part2(scanner *bufio.Scanner) string {
	return strconv.Itoa(solver.NewEquationSolver().SolveEquations(scanner, true))
}
