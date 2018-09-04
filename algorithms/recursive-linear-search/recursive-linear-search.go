package recursivelinearsearch

// RecursiveLinearSearch starts at the beginning of array A and looks for the value x.
// It returns the index of x and a boolean indicating whether the value x was
// found in the array. The value of the returned index is undefined if x is
// not found in the array.
func RecursiveLinearSearch(A []int, i, x int) (int, bool) {
	if i > len(A)-1 {
		return 0, false
	}

	if A[i] == x {
		return i, true
	}

	return RecursiveLinearSearch(A, i+1, x)
}

// RecursiveLinearSearchAlt implements RecursiveLinearSearch without the need to pass
// in the additional index as part of the function call.
func RecursiveLinearSearchAlt(A []int, x int) (int, bool) {
	if len(A) == 0 {
		return 0, false
	}

	if A[0] == x {
		return 0, true
	}

	l, ok := RecursiveLinearSearchAlt(A[1:], x)
	if ok == false {
		return 0, false
	}
	return l + 1, true
}
