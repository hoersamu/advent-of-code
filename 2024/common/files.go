package common

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

func ScanWithDelimitersAsInt(scanner *bufio.Scanner, delimiter string) [][]int {
	return scanWithDelimiters(scanner, delimiter, func(s string) int {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal("Unable to parse string to int", err)
		}
		return num
	})
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
		row := strings.Split(text, delimiter)
		transformed := []T{}
		for _, cell := range row {
			transformed = append(transformed, transformer(cell))
		}
		rows = append(rows, transformed)
	}

	return rows
}

func ScanToString(scanner *bufio.Scanner) string {
	builder := strings.Builder{}
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}
	return builder.String()
}