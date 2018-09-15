package countingsort

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var testCases = []struct {
	Name  string
	A     []int
	m     int
	equal []int
	less  []int
	B     []int
}{
	{
		Name:  "at least one of each value",
		A:     []int{0, 1, 0, 3, 2, 1, 1, 1, 4, 3},
		m:     5,
		equal: []int{2, 4, 1, 2, 1},
		less:  []int{0, 2, 6, 7, 9},
		B:     []int{0, 0, 1, 1, 1, 1, 2, 3, 3, 4},
	},
	{
		Name:  "some values not present",
		A:     []int{0, 1, 0, 3, 2, 1, 1, 1, 5, 3},
		m:     6,
		equal: []int{2, 4, 1, 2, 0, 1},
		less:  []int{0, 2, 6, 7, 9, 9},
		B:     []int{0, 0, 1, 1, 1, 1, 2, 3, 3, 5},
	},
}

func TestCountKeysEqual(t *testing.T) {
	for _, tc := range testCases {
		equal := countKeysEqual(tc.A, tc.m)

		if got, want := equal, tc.equal; reflect.DeepEqual(got, want) != true {
			t.Errorf("\n got: %v\nwant: %v", got, want)
		}
	}
}

func TestCountKeysLess(t *testing.T) {
	for _, tc := range testCases {
		less := countKeysLess(tc.equal)

		if got, want := less, tc.less; reflect.DeepEqual(got, want) != true {
			t.Errorf("\n got: %v\nwant: %v", got, want)
		}
	}
}

func TestRearrange(t *testing.T) {
	for _, tc := range testCases {
		B := rearrange(tc.A, tc.less)

		if got, want := B, tc.B; reflect.DeepEqual(got, want) != true {
			t.Errorf("\n got: %v\nwant: %v", got, want)
		}
	}
}

func TestCountingSort(t *testing.T) {
	for _, tc := range testCases {
		B := CountingSort(tc.A, tc.m)

		if got, want := B, tc.B; reflect.DeepEqual(got, want) != true {
			t.Errorf("\n got: %v\nwant: %v", got, want)
		}
	}
}

func benchmarkCountingSort(i, m int, b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	A := make([]int, i)
	a := make([]int, i)
	for k := range A {
		A[k] = rand.Intn(m)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		copy(a, A)
		CountingSort(a, m)
	}
}

func BenchmarkCountingSortArraySize(b *testing.B) {

	m := 10

	for i := 1000; i <= 10000; i = i + 1000 {
		b.Run(fmt.Sprintf("_%d", i), func(b *testing.B) {
			benchmarkCountingSort(i, m, b)
		})
	}
}

func BenchmarkCountingSortKeySize(b *testing.B) {

	i := 1000
	for m := 1000; m <= 10000; m = m + 1000 {
		b.Run(fmt.Sprintf("_%d", m), func(b *testing.B) {
			benchmarkCountingSort(i, m, b)
		})
	}
}
