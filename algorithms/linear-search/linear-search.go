package linearsearch

// LinearSearch starts at the beginning of array A and looks for the value x.
// It returns the index of x and a boolean indicating whether the value x was
// found in the array. The value of the returned index is undefined if x is
// not found in the array.
func LinearSearch(A []int, x int) (int, bool) {
	index := 0
	found := false

	for i := range A {
		if A[i] == x {
			index, found = i, true
		}
	}

	return index, found
}
