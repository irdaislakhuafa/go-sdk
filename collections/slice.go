package collections

import "slices"

// Will check is elements of 2 slice is equals with ignored order
func IsElementsEquals[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for _, v := range b {
		if !slices.Contains(a, v) {
			return false
		}
	}
	return true
}
