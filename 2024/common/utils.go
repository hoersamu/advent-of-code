package common

import "math"

func RemoveIndex(slice []int, s int) []int {
	result := make([]int, 0, len(slice)-1)
	result = append(result, slice[:s]...)
	result = append(result, slice[s+1:]...)
	return result
}

func GetDistance(a int, b int) int {
	return int(math.Abs(float64(a - b)))
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
