package mergesort

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

func TestMergeSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			MergeSort(tc.A)
			if got, want := tc.A, tc.S; reflect.DeepEqual(got, want) != true {
				t.Errorf("unexpected array returned:\n\t got:%v\n\twant:%v\n", got, want)
			}
		})
	}
}

func benchmarkMergeSort(i int, b *testing.B) {
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
		MergeSort(a)
	}
}

func BenchmarkMergeSort_1000(b *testing.B)  { benchmarkMergeSort(1000, b) }
func BenchmarkMergeSort_2000(b *testing.B)  { benchmarkMergeSort(2000, b) }
func BenchmarkMergeSort_3000(b *testing.B)  { benchmarkMergeSort(3000, b) }
func BenchmarkMergeSort_4000(b *testing.B)  { benchmarkMergeSort(4000, b) }
func BenchmarkMergeSort_5000(b *testing.B)  { benchmarkMergeSort(5000, b) }
func BenchmarkMergeSort_6000(b *testing.B)  { benchmarkMergeSort(6000, b) }
func BenchmarkMergeSort_7000(b *testing.B)  { benchmarkMergeSort(7000, b) }
func BenchmarkMergeSort_8000(b *testing.B)  { benchmarkMergeSort(8000, b) }
func BenchmarkMergeSort_9000(b *testing.B)  { benchmarkMergeSort(9000, b) }
func BenchmarkMergeSort_10000(b *testing.B) { benchmarkMergeSort(10000, b) }
