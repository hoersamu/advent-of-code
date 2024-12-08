package common

import "math"

// Returns the median element of a slice
func GetMedianElement[T any](slice []T) T {
	return slice[int(math.Floor(float64(len(slice))/2))]
}
