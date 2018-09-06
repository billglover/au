package selectionsort

// SelectionSort takes an array of int and sorts it using the Selection Sort algorithm.
func SelectionSort(A []int) {

	for i := range A {
		smallest := i
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[smallest] {
				smallest = j
			}
		}
		A[i], A[smallest] = A[smallest], A[i]
	}

}
