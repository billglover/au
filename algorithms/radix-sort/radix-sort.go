package radixsort

// RadixSort takes a slice of strings of length m and returns a new slice
// containing the elements of A sorted in ascending order.
func RadixSort(A []string, m int) []string {
	B := make([]string, len(A))
	copy(B, A)

	for c := m - 1; c >= 0; c-- {
		B = countingSort(B, c)
	}
	return B
}

func countingSort(A []string, c int) []string {
	equal := make([]int, 36)
	for i := range A {

		// use 0-9 for digits
		// use 10-35 for letters
		key := A[i][c]
		if key >= 65 {
			key = key - 7
		}
		key = key - 48
		equal[key]++
	}

	less := make([]int, 36)
	for j := range equal {
		if j == 0 {
			continue
		}
		less[j] = less[j-1] + equal[j-1]
	}

	B := make([]string, len(A))
	next := make([]int, 36)

	copy(next, less)

	for i := range A {
		key := A[i][c]
		if key >= 65 {
			key = key - 7
		}
		key = key - 48
		equal[key]++

		index := next[key]
		B[index] = A[i]
		next[key]++
	}

	return B
}
