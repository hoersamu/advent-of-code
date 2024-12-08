package main

import (
	"bufio"
	"common"
)

func readInput(scanner *bufio.Scanner) ([][]int, [][]int) {
	var rules [][]int
	var updates [][]int

	// Read rules until empty line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := common.SplitAndTransform(line, "|", common.MustAtoi)
		rules = append(rules, parts[:2])
	}

	// Read updates
	for scanner.Scan() {
		line := scanner.Text()
		parts := common.SplitAndTransform(line, ",", common.MustAtoi)
		updates = append(updates, parts)
	}

	return rules, updates
}

func buildStateMachine(rules [][]int) map[int]map[int]bool {
	// For each number, store what numbers must NOT come before it
	stateMachine := make(map[int]map[int]bool)
	for _, rule := range rules {
		before, after := rule[0], rule[1]
		// If 'after' appears, 'before' must not come later
		if stateMachine[after] == nil {
			stateMachine[after] = make(map[int]bool)
		}
		stateMachine[after][before] = true
	}
	return stateMachine
}

func verifySequence(sequence []int, stateMachine map[int]map[int]bool) bool {
	// Keep track of position of each number
	positions := make(map[int]int)
	for pos, num := range sequence {
		positions[num] = pos
	}

	// Check each number's position against its requirements
	for pos, num := range sequence {
		if requirements, exists := stateMachine[num]; exists {
			for mustBeBefore := range requirements {
				// Only check if both numbers exist in the sequence
				if beforePos, exists := positions[mustBeBefore]; exists {
					if beforePos > pos {
						return false
					}
				}
			}
		}
	}
	return true
}

func main() {
	common.Solve(Part1)
	common.Solve(Part2)
}

func Part1(scanner *bufio.Scanner) int {
	rules, updates := readInput(scanner)
	stateMachine := buildStateMachine(rules)

	result := 0
	for _, update := range updates {
		if verifySequence(update, stateMachine) {
			result += common.GetMedianElement(update)
		}
	}

	return result
}

func Part2(scanner *bufio.Scanner) int {
	rules, updates := readInput(scanner)
	stateMachine := buildStateMachine(rules)

	result := 0
	for _, update := range updates {
		if !verifySequence(update, stateMachine) {
			// Fix order of update to pass state machine
			sorted := make([]int, len(update))
			copy(sorted, update)
			for i := 0; i < len(sorted)-1; i++ {
				for j := 0; j < len(sorted)-i-1; j++ {
					// If these numbers have a rule and are in wrong order, swap them
					if requirements, exists := stateMachine[sorted[j+1]]; exists {
						if requirements[sorted[j]] {
							sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
						}
					}
				}
			}
			update = sorted
			result += common.GetMedianElement(update)
		}
	}

	return result
}
