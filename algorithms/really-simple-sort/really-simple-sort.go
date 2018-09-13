package reallysimplesort

// ReallySimpleSort takes an array of int and sorts it using the Really Simple Sort algorithm.
// It assumes that all values in the array, A, are either 1 or 2. If any other values are
// passed the result is undefined.
func ReallySimpleSort(A []int) []int {
	k := 0
	for i := range A {
		if A[i] == 1 {
			k++
		}
	}

	for i := range A {
		switch i < k {
		case true:
			A[i] = 1
		case false:
			A[i] = 2
		}
	}

	return A
}
