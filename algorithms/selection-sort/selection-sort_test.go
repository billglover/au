package selectionsort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var testCases = []struct {
	Name string
	A    []int
	S    []int
}{
	{
		Name: "array already sorted",
		A:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		S:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		Name: "unsorted positive integers",
		A:    []int{1, 7, 3, 0, 9, 4, 8, 2, 6, 5},
		S:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		Name: "unsorted positive and negative integers",
		A:    []int{-2, 4, 1, 0, -1, -4, 3, 2, 5, -3},
		S:    []int{-4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
	},
}

func TestLinearSearch(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			SelectionSort(tc.A)
			if got, want := tc.A, tc.S; reflect.DeepEqual(got, want) != true {
				t.Errorf("unexpected array returned:\n\t got:%v\n\twant:%v\n", got, want)
			}
		})
	}
}

func benchmarkSelectionSort(i int, b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	A := rand.Perm(i)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		SelectionSort(A)
	}
}

func BenchmarkSelectionSort_1000(b *testing.B)  { benchmarkSelectionSort(1000, b) }
func BenchmarkSelectionSort_2000(b *testing.B)  { benchmarkSelectionSort(2000, b) }
func BenchmarkSelectionSort_3000(b *testing.B)  { benchmarkSelectionSort(3000, b) }
func BenchmarkSelectionSort_4000(b *testing.B)  { benchmarkSelectionSort(4000, b) }
func BenchmarkSelectionSort_5000(b *testing.B)  { benchmarkSelectionSort(5000, b) }
func BenchmarkSelectionSort_6000(b *testing.B)  { benchmarkSelectionSort(6000, b) }
func BenchmarkSelectionSort_7000(b *testing.B)  { benchmarkSelectionSort(7000, b) }
func BenchmarkSelectionSort_8000(b *testing.B)  { benchmarkSelectionSort(8000, b) }
func BenchmarkSelectionSort_9000(b *testing.B)  { benchmarkSelectionSort(9000, b) }
func BenchmarkSelectionSort_10000(b *testing.B) { benchmarkSelectionSort(10000, b) }
