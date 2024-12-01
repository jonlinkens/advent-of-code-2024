package util

import "slices"

func RemoveFromSlice[T comparable](slice []T, val T) []T {

	index := slices.Index(slice, val)

	lastIndex := len(slice) - 1

	slice[index] = slice[lastIndex]

	return slice[:lastIndex]

}
