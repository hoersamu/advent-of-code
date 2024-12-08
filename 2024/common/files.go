package common

import (
	"bufio"
	"strings"
)

func ScanWithDelimitersAsInt(scanner *bufio.Scanner, delimiter string) [][]int {
	return scanWithDelimiters(scanner, delimiter, MustAtoi)
}

func ScanWithDelimiters(scanner *bufio.Scanner, delimiter string) [][]string {
	return scanWithDelimiters(scanner, delimiter, func(s string) string { return s })
}

func scanWithDelimiters[T any](
	scanner *bufio.Scanner,
	delimiter string,
	transformer func(string) T,
) [][]T {
	var rows [][]T

	for scanner.Scan() {
		text := scanner.Text()
		rows = append(rows, SplitAndTransform(text, delimiter, transformer))
	}

	return rows
}

func SplitAndTransform[T any](input string, delimiter string, transformer func(string) T) []T {
	result := []T{}
	for _, s := range strings.Split(input, delimiter) {
		result = append(result, transformer(s))
	}
	return result
}

func ScanToString(scanner *bufio.Scanner) string {
	builder := strings.Builder{}
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}
	return builder.String()
}

func ScanToLines(scanner *bufio.Scanner) []string {
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
