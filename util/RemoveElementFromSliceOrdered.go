package util

func RemoveFromSliceOrdered[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}

	newSlice := make([]T, len(slice)-1)

	copy(newSlice, slice[:index])

	copy(newSlice[index:], slice[index+1:])

	return newSlice
}
