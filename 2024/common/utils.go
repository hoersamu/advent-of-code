package common

import (
	"math"
	"strconv"
	"strings"
)

// Returns the absolute distance between two integers
func GetDistance(a int, b int) int {
	return int(math.Abs(float64(a - b)))
}

// Reverses a string
func ReverseString(s string) string {
	runes := []rune(s)
	sb := strings.Builder{}
	for i := len(runes) - 1; i >= 0; i-- {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}

// Converts a string to an integer, panics if the conversion fails
func MustAtoi(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}
