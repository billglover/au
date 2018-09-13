package reallysimplesort

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

var testCases = []struct {
	Name string
	A    []int
	S    []int
}{
	{
		Name: "array already sorted",
		A:    []int{1, 1, 1, 1, 1, 1, 1, 2, 2, 2},
		S:    []int{1, 1, 1, 1, 1, 1, 1, 2, 2, 2},
	},
	{
		Name: "array is unsorted",
		A:    []int{2, 1, 1, 1, 1, 1, 1, 2, 2, 1},
		S:    []int{1, 1, 1, 1, 1, 1, 1, 2, 2, 2},
	},
}

func TestReallySimpleSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ReallySimpleSort(tc.A)
			if got, want := tc.A, tc.S; reflect.DeepEqual(got, want) != true {
				t.Errorf("unexpected array returned:\n\t got:%v\n\twant:%v\n", got, want)
			}
		})
	}
}

func benchmarkReallySimpleSort(i int, b *testing.B) {
	A := make([]int, i)
	a := make([]int, i)

	for k := range A {
		A[k] = rand.Intn(1) + 1
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		copy(a, A) // assumed to be Θ(n)
		ReallySimpleSort(a)
	}
}

func BenchmarkReallySimpleSort(b *testing.B) {

	for i := 1000; i <= 10000; i = i + 1000 {
		b.Run(fmt.Sprintf("ArraySize_%d", i), func(b *testing.B) {
			benchmarkReallySimpleSort(i, b)
		})
	}
}
