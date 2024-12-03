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
