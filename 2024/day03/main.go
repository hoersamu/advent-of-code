package main

import (
	"bufio"
	"common"
	"regexp"
	"strconv"
	"strings"
)

func processMul(text string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(text, -1)

	sum := 0

	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		product := num1 * num2
		sum += product
	}

	return sum
}

func main() {
	common.Solve(Part1)
	common.Solve(Part2)
}

func Part1(scanner *bufio.Scanner) string {
	return strconv.Itoa(processMul(common.ScanToString(scanner)))
}

func Part2(scanner *bufio.Scanner) string {
	text := common.ScanToString(scanner)

	parts := strings.Split(text, "do()")
	builder := strings.Builder{}
	for _, part := range parts {
		builder.WriteString(strings.Split(part, "don't()")[0])
	}

	return strconv.Itoa(processMul(builder.String()))
}
