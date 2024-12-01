package util

func AbsoluteDiff(a int, b int) int {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff
}
