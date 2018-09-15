package countingsort

// CountingSort takes an array of ints in the range 0 to m-1 and returns a new array
// containing the elements of A sorted in ascending order.
func CountingSort(A []int, m int) []int {
	equal := countKeysEqual(A, m)
	less := countKeysLess(equal)
	B := rearrange(A, less)
	return B
}

// CountKeysEqual takes a slice, A,  of ints and the number of unique values in A, m.
// It returns a slice such that each element in the slice contains the number of
// elements in A that equal the element index.
//
// It assumes that there are no elements present in A that lie outside the range 0 to m-1.
// If integers are outside that range, the function will panic.
func countKeysEqual(A []int, m int) []int {
	equal := make([]int, m)
	for i := range A {
		key := A[i]
		equal[key]++
	}
	return equal
}

// CountKeysLess takes an slice of key counts (ints) and returns an slice of equal length
// that contains the cumulative total of values in the input array.
func countKeysLess(equal []int) []int {
	less := make([]int, len(equal))
	for j := range equal {
		if j == 0 {
			continue
		}
		less[j] = less[j-1] + equal[j-1]
	}
	return less
}

// Rearrange takes a slice of ints, A along with a cumulative count of
// the values in A and returns a new sorted slice containing the
// elements of A.
func rearrange(A, less []int) []int {
	B := make([]int, len(A))
	next := make([]int, len(less))

	copy(next, less)

	for i := range A {
		key := A[i]
		index := next[key]
		B[index] = A[i]
		next[key]++
	}

	return B
}
