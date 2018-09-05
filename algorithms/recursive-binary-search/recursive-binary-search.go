package recursivebinarysearch

// BinarySearch takes a sorted array A and looks for the value x.
// It returns the index of x and a boolean indicating whether the value x was
// found in the array. The value of the returned index is undefined if x is
// not found in the array.
func BinarySearch(A []int, x int) (int, bool) {
	return recursiveBinarySearch(A, 0, len(A)-1, x)
}

func recursiveBinarySearch(A []int, p, r, x int) (int, bool) {
	if p > r {
		return 0, false
	}

	q := (p + r) / 2

	switch {

	case A[q] == x:
		return q, true

	case A[q] > x:
		return recursiveBinarySearch(A, p, q-1, x)

	default:
		return recursiveBinarySearch(A, q+1, r, x)
	}
}
