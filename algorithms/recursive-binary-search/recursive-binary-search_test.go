package recursivebinarysearch

import (
	"math/rand"
	"testing"
	"time"
)

var testCases = []struct {
	Name  string
	A     []int
	x     int
	i     int
	found bool
}{
	{
		Name:  "value present in array",
		A:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		x:     3,
		i:     3,
		found: true,
	},
	{
		Name:  "value not present in array",
		A:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		x:     10,
		i:     0,
		found: false,
	},
}

func TestRecursiveBinarySearch(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			loc, found := BinarySearch(tc.A, tc.x)

			if got, want := found, tc.found; got != want {
				t.Errorf("unexpected found: got %v, want %v", got, want)
			}

			if got, want := loc, tc.i; got != want {
				t.Errorf("unexpected index: got %d, want %d", got, want)
			}
		})
	}

}

func benchmarkRecursiveBinarySearch(i int, b *testing.B) {
	var l int
	var f bool

	rand.Seed(time.Now().UnixNano())
	A := make([]int, i)
	for k := range A {
		A[k] = k
	}
	x := -1

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		l, f = BinarySearch(A, x)
	}
	loc, found = l, f
}

func BenchmarkRecursiveBinarySearch_100(b *testing.B)    { benchmarkRecursiveBinarySearch(100, b) }
func BenchmarkRecursiveBinarySearch_200(b *testing.B)    { benchmarkRecursiveBinarySearch(200, b) }
func BenchmarkRecursiveBinarySearch_400(b *testing.B)    { benchmarkRecursiveBinarySearch(400, b) }
func BenchmarkRecursiveBinarySearch_800(b *testing.B)    { benchmarkRecursiveBinarySearch(800, b) }
func BenchmarkRecursiveBinarySearch_1600(b *testing.B)   { benchmarkRecursiveBinarySearch(1600, b) }
func BenchmarkRecursiveBinarySearch_3200(b *testing.B)   { benchmarkRecursiveBinarySearch(3200, b) }
func BenchmarkRecursiveBinarySearch_6400(b *testing.B)   { benchmarkRecursiveBinarySearch(6400, b) }
func BenchmarkRecursiveBinarySearch_12800(b *testing.B)  { benchmarkRecursiveBinarySearch(12800, b) }
func BenchmarkRecursiveBinarySearch_25600(b *testing.B)  { benchmarkRecursiveBinarySearch(25600, b) }
func BenchmarkRecursiveBinarySearch_51200(b *testing.B)  { benchmarkRecursiveBinarySearch(51200, b) }
func BenchmarkRecursiveBinarySearch_102400(b *testing.B) { benchmarkRecursiveBinarySearch(102400, b) }

var loc int
var found bool
