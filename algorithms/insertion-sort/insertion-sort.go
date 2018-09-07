package insertionsort

// InsertionSort takes an array of int and sorts it using the Selection Sort algorithm.
func InsertionSort(A []int) {
	for i := 1; i < len(A); i++ {
		key := A[i]
		j := i - 1

		for j >= 0 && A[j] > key {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = key
	}
}
