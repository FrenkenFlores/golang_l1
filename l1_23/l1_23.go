package l1_23

import "slices"

func DeleteSliceElements(s []int, i, j int) []int {
	return slices.Delete(s, i, j)
}

func DeleteSliceElement(s []int, i int) []int {
	if i < 0 || i >= len(s) {
		return s
	}

	newSlice := make([]int, len(s)-1)
	copy(newSlice[:i], s[:i])
	copy(newSlice[i:], s[i+1:])
	return newSlice
}
