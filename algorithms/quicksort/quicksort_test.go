package quicksort

import (
	"fmt"
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

func TestQuickSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			QuickSort(tc.A)
			if got, want := tc.A, tc.S; reflect.DeepEqual(got, want) != true {
				t.Errorf("unexpected array returned:\n\t got:%v\n\twant:%v\n", got, want)
			}
		})
	}
}

func benchmarkQuickSort(i int, b *testing.B) {
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
		QuickSort(a)
	}
}

func BenchmarkQuickSort(b *testing.B) {

	for i := 1000; i <= 10000; i = i + 1000 {
		b.Run(fmt.Sprintf("ArraySize_%d", i), func(b *testing.B) {
			benchmarkQuickSort(i, b)
		})
	}
}
