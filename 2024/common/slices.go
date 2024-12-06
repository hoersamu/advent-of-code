package common

import "math"

func GetMedianElement[T any](slice []T) T {
	return slice[int(math.Floor(float64(len(slice))/2))]
}
