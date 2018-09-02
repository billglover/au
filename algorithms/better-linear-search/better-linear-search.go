package betterlinearsearch

// BetterLinearSearch starts at the beginning of array A and looks for the
// value x. Unlike LinearSearch, it returns as soon as it has found a result.
// It returns the index of x and a boolean indicating whether the value x was
// found in the array. The value of the returned index is undefined if x is
// not found in the array.
func LinearSearch(A []int, x int) (int, bool) {
	for i := range A {
		if A[i] == x {
			return i, true
		}
	}

	return 0, false
}
