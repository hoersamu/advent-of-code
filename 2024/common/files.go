package common

import (
	"bufio"
	"strings"
)

// Scans the input into a list of rows, where each row is a list of ints
func ScanWithDelimitersAsInt(scanner *bufio.Scanner, delimiter string) [][]int {
	return scanWithDelimiters(scanner, delimiter, MustAtoi)
}

// Scans the input into a list of rows, where each row is a list of strings
func ScanWithDelimiters(scanner *bufio.Scanner, delimiter string) [][]string {
	return scanWithDelimiters(scanner, delimiter, id[string])
}

// Generic utility function to scan the input into a list of rows,
// where each row is a list of transformed strings
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

// Splits the input into a list of strings using the delimiter and transforms each string
func SplitAndTransform[T any](input string, delimiter string, transformer func(string) T) []T {
	result := []T{}
	for _, s := range strings.Split(input, delimiter) {
		result = append(result, transformer(s))
	}
	return result
}

// Scans the input into a single string
func ScanToString(scanner *bufio.Scanner) string {
	builder := strings.Builder{}
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}
	return builder.String()
}

// Scans the input into a list of lines
func ScanToLines(scanner *bufio.Scanner) []string {
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// Identity function -> returns the input
func id[T any](x T) T { return x }
