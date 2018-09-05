package binarysearch

// BinarySearch takes a sorted array A and looks for the value x.
// It returns the index of x and a boolean indicating whether the value x was
// found in the array. The value of the returned index is undefined if x is
// not found in the array.
func BinarySearch(A []int, x int) (int, bool) {

	p, r := 1, len(A)-1

	for p <= r {
		q := (p + r) / 2

		switch {

		case A[q] == x:
			return q, true

		case A[q] > x:
			r = q - 1

		default:
			p = q + 1

		}
	}

	return 0, false
}
