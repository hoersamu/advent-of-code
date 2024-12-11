package main

import (
	"bufio"
	"common"
)

func main() {
	common.Solve(Part1)
	common.Solve(Part2)
}

/**
 * The Order of the Numbers is actually irrelevant since we are not joining numbers only splitting and modifying them
 * Instead of using a real slice we use a map to track the number of times each number appears
 * This is a lot faster than using a slice and allows us to split the numbers without having to worry about the order
 */
func runCode(scanner *bufio.Scanner, iterations int) int {
	const multiplier = 2024
	scanner.Scan()
	numbers := common.SplitAndTransform(scanner.Text(), " ", common.MustAtoi)

	numCount := make(map[int]int)
	for _, n := range numbers {
		numCount[n]++
	}

	for i := 0; i < iterations; i++ {
		newCount := make(map[int]int)
		for num, count := range numCount {
			if num == 0 {
				newCount[1] += count
				continue
			}

			digits := common.GetNumberOfDigits(num)
			if common.IsEven(digits) {
				divider := 1
				for j := 0; j < digits/2; j++ {
					divider *= 10
				}
				newCount[num/divider] += count
				newCount[num%divider] += count
			} else {
				newCount[num*multiplier] += count
			}
		}
		numCount = newCount
	}

	total := 0
	for _, count := range numCount {
		total += count
	}
	return total
}

func Part1(scanner *bufio.Scanner) int {
	return runCode(scanner, 25)
}

func Part2(scanner *bufio.Scanner) int {
	return runCode(scanner, 75)
}
