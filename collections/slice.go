package collections

import "slices"

func IsElementsEquals[T comparable](a, b []T) bool {
	for _, v := range a {
		if !slices.Contains(b, v) {
			return false
		}
	}
	return true
}
