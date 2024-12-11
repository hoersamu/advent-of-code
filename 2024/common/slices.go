package common

import (
	"math"
	"slices"
)

// Returns the median element of a slice
func GetMedianElement[T any](slice []T) T {
	return slice[int(math.Floor(float64(len(slice))/2))]
}

func DeleteElement[T any](slice []T, index int) ([]T, T) {
	element := slice[index]
	return slices.Delete(slice, index, index+1), element
}
