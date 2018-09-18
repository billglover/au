package radixsort

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

var testCases = []struct {
	Name string
	A    []string
	m    int
	B    []string
}{
	{
		Name: "unsorted arrayr",
		A: []string{
			"PL4ZQ2", "KI4WR2", "JL2ZV3", "PY2ZR5", "XI7FS6", "XL8FQ6", "JI8FR9", "KV7WS9",
		},
		m: 6,
		B: []string{
			"JI8FR9", "JL2ZV3", "KI4WR2", "KV7WS9", "PL4ZQ2", "PY2ZR5", "XI7FS6", "XL8FQ6",
		},
	},
	{
		Name: "sorted array",
		A: []string{
			"JI8FR9", "JL2ZV3", "KI4WR2", "KV7WS9", "PL4ZQ2", "PY2ZR5", "XI7FS6", "XL8FQ6",
		},
		m: 6,
		B: []string{
			"JI8FR9", "JL2ZV3", "KI4WR2", "KV7WS9", "PL4ZQ2", "PY2ZR5", "XI7FS6", "XL8FQ6",
		},
	},
}

func TestRadixSort(t *testing.T) {
	for _, tc := range testCases {
		B := RadixSort(tc.A, tc.m)

		if got, want := B, tc.B; reflect.DeepEqual(got, want) != true {
			t.Errorf("\n got: %v\nwant: %v", got, want)
		}
	}
}

func benchmarkRadixSort(i, m int, b *testing.B) {
	A := make([]string, i)
	a := make([]string, i)
	for k := range A {
		A[k] = getKey(m)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		copy(a, A)
		RadixSort(a, m)
	}
}

func BenchmarkRadixSort(b *testing.B) {

	m := 6

	for i := 1000; i <= 10000; i = i + 1000 {
		b.Run(fmt.Sprintf("_%d", i), func(b *testing.B) {
			benchmarkRadixSort(i, m, b)
		})
	}
}

var chars = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func getKey(size int) string {
	l := len(chars)
	k := make([]rune, size)
	for i := range k {
		k[i] = chars[rand.Intn(l)]
	}
	return string(k)
}
