package common

import "math"

func RemoveIndex[T any](slice []T, s int) []T {
	result := make([]T, 0, len(slice)-1)
	result = append(result, slice[:s]...)
	result = append(result, slice[s+1:]...)
	return result
}

func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func IndexOf[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func GetMedianElement[T any](slice []T) T {
	return slice[int(math.Floor(float64(len(slice))/2))]
}
