package util

// RemoveFromArrayStable removes the first occurrence of `match` in the given `array`.
// It returns True if one element was removed. False otherwise.
// The order of the remaining elements is not changed.
func RemoveFromArrayStable(array []uint64, match uint64) ([]uint64, bool) {
	for i, other := range array {
		if other == match {
			return append(array[0:i], array[i+1:]...), true
		}
	}
	return array, false
}

func Contains(array []uint64, match uint64) bool {
	for _, other := range array {
		if other == match {
			return true
		}
	}
	return false
}
