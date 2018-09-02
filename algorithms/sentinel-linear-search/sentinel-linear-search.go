package sentinellinearsearch

// SentinelLinearSearch starts at the beginning of array A and looks for the
// value x. It offers a slight improvement in run-time over BetterLinearSearch.
// It returns the index of x and a boolean indicating whether the value x was
// found in the array. The value of the returned index is undefined if x is
// not found in the array.
func SentinelLinearSearch(A []int, x int) (int, bool) {
	n := len(A) - 1
	last := A[n]
	A[n] = x
	i := 0

	for {
		if A[i] == x {
			break
		}
		i++
	}

	A[n] = last
	if i < n || A[n] == x {
		return i, true
	}

	return 0, false
}
