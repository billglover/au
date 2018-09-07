package insertionsort

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

func TestInsertionSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			InsertionSort(tc.A)
			if got, want := tc.A, tc.S; reflect.DeepEqual(got, want) != true {
				t.Errorf("unexpected array returned:\n\t got:%v\n\twant:%v\n", got, want)
			}
		})
	}
}

func benchmarkInsertionSort(i int, b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	A := make([]int, i)
	a := make([]int, i)
	for k := range A {
		A[k] = i - k
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		copy(a, A)
		b.StartTimer()
		InsertionSort(a)
	}
}

func BenchmarkInsertionSort_1000(b *testing.B)  { benchmarkInsertionSort(1000, b) }
func BenchmarkInsertionSort_2000(b *testing.B)  { benchmarkInsertionSort(2000, b) }
func BenchmarkInsertionSort_3000(b *testing.B)  { benchmarkInsertionSort(3000, b) }
func BenchmarkInsertionSort_4000(b *testing.B)  { benchmarkInsertionSort(4000, b) }
func BenchmarkInsertionSort_5000(b *testing.B)  { benchmarkInsertionSort(5000, b) }
func BenchmarkInsertionSort_6000(b *testing.B)  { benchmarkInsertionSort(6000, b) }
func BenchmarkInsertionSort_7000(b *testing.B)  { benchmarkInsertionSort(7000, b) }
func BenchmarkInsertionSort_8000(b *testing.B)  { benchmarkInsertionSort(8000, b) }
func BenchmarkInsertionSort_9000(b *testing.B)  { benchmarkInsertionSort(9000, b) }
func BenchmarkInsertionSort_10000(b *testing.B) { benchmarkInsertionSort(10000, b) }
