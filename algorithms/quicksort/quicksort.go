package quicksort

import "math/rand"

// QuickSort takes an array of int and sorts it using the Quick Sort algorithm.
func QuickSort(A []int) {
	quickSort(A, 0, len(A)-1)
}

func quickSort(A []int, p, r int) {
	if p >= r {
		return
	}

	n := rand.Intn(r-p) + p
	A[r], A[n] = A[n], A[r]

	q := partition(A, p, r)
	quickSort(A, p, q-1)
	quickSort(A, q+1, r)

}

func partition(A []int, p, r int) int {
	q := p
	for u := p; u < r; u++ {
		if A[u] <= A[r] {
			A[q], A[u] = A[u], A[q]
			q++
		}
	}
	A[q], A[r] = A[r], A[q]
	return q
}
