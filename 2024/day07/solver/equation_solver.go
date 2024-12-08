package solver

import (
	"bufio"
	"common"
	"fmt"
	"strings"
)

type Operation interface {
	Apply(a, b int) int
}

type Equation struct {
	result  int
	numbers []int
}

type AddOperation struct{}
type MultiplyOperation struct{}
type ConcatOperation struct{}

func (o AddOperation) Apply(a, b int) int      { return a + b }
func (o MultiplyOperation) Apply(a, b int) int { return a * b }
func (o ConcatOperation) Apply(a, b int) int {
	return common.MustAtoi(fmt.Sprintf("%d%d", a, b))
}

type EquationSolver struct {
	operations []Operation
}

func NewEquationSolver() *EquationSolver {
	return &EquationSolver{
		operations: []Operation{
			AddOperation{},
			MultiplyOperation{},
		},
	}
}

func (s *EquationSolver) withConcat() *EquationSolver {
	s.operations = append(s.operations, ConcatOperation{})
	return s
}

func (s *EquationSolver) isSolvable(target int, numbers []int) bool {
	if len(numbers) == 2 {
		return s.checkOperations(target, numbers[0], numbers[1])
	}

	for _, op := range s.operations {
		result := op.Apply(numbers[0], numbers[1])
		remaining := append([]int{result}, numbers[2:]...)
		if s.isSolvable(target, remaining) {
			return true
		}
	}
	return false
}

func (s *EquationSolver) checkOperations(target, a, b int) bool {
	for _, op := range s.operations {
		if op.Apply(a, b) == target {
			return true
		}
	}
	return false
}

func (s *EquationSolver) SolveEquations(scanner *bufio.Scanner, includeConcat bool) int {
	if includeConcat {
		s.withConcat()
	}

	equations := parseInput(scanner)
	total := 0

	for _, eq := range equations {
		if s.isSolvable(eq.result, eq.numbers) {
			total += eq.result
		}
	}
	return total
}

func parseInput(scanner *bufio.Scanner) []Equation {
	lines := common.ScanToLines(scanner)
	equations := []Equation{}
	for _, line := range lines {
		resultStr, numbersStr := strings.Split(line, ": ")[0], strings.Split(line, ": ")[1]
		result := common.MustAtoi(resultStr)
		numbers := common.SplitAndTransform(numbersStr, " ", common.MustAtoi)

		equations = append(equations, Equation{result, numbers})
	}
	return equations
}
