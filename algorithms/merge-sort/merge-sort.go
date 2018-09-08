package mergesort

import (
	"math"
)

// MergeSort takes an array of int and sorts it using the Merge Sort algorithm.
func MergeSort(A []int) {
	if len(A) == 1 {
		return
	}

	q := len(A) / 2

	MergeSort(A[:q])
	MergeSort(A[q:])
	Merge(A, q)
}

// Merge takes an array and index that identifies the boundary of two subarrays,
// A[:q] and A[q:]. It merges these two subarrays to ensure that the
// elements of `A[:] are sorted.`
func Merge(A []int, q int) {
	n1 := len(A[:q])
	n2 := len(A[q:])

	B := make([]int, n1+1)
	C := make([]int, n2+1)

	copy(B, A[:q])
	copy(C, A[q:])

	// We set an additional final element in each sub-array to be the
	// maximum allowed value to ensure we don't compare with an element
	// that exists outside the array bounds.
	B[n1] = math.MaxInt64
	C[n2] = math.MaxInt64

	i, j := 0, 0
	for k := range A {
		if B[i] <= C[j] {
			A[k] = B[i]
			i++
		} else {
			A[k] = C[j]
			j++
		}
	}
}
