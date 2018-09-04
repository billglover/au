package recursivelinearsearch

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

func TestLinearSearch(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			loc, found := RecursiveLinearSearch(tc.A, 0, tc.x)

			if got, want := found, tc.found; got != want {
				t.Errorf("unexpected found: got %v, want %v", got, want)
			}

			if got, want := loc, tc.i; got != want {
				t.Errorf("unexpected index: got %d, want %d", got, want)
			}
		})
	}
}

func TestAltLinearSearch(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			loc, found := RecursiveLinearSearchAlt(tc.A, tc.x)

			if got, want := found, tc.found; got != want {
				t.Errorf("unexpected found: got %v, want %v", got, want)
			}

			if got, want := loc, tc.i; got != want {
				t.Errorf("unexpected index: got %d, want %d", got, want)
			}
		})
	}
}

func benchmarkRecursiveLinearSearch(i int, b *testing.B) {
	var l int
	var f bool

	rand.Seed(time.Now().UnixNano())
	A := rand.Perm(i)
	x := -1

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		l, f = RecursiveLinearSearch(A, 0, x)
	}
	loc, found = l, f
}

func BenchmarkRecursiveLinearSearch_100(b *testing.B)    { benchmarkRecursiveLinearSearch(100, b) }
func BenchmarkRecursiveLinearSearch_200(b *testing.B)    { benchmarkRecursiveLinearSearch(200, b) }
func BenchmarkRecursiveLinearSearch_400(b *testing.B)    { benchmarkRecursiveLinearSearch(400, b) }
func BenchmarkRecursiveLinearSearch_800(b *testing.B)    { benchmarkRecursiveLinearSearch(800, b) }
func BenchmarkRecursiveLinearSearch_1600(b *testing.B)   { benchmarkRecursiveLinearSearch(1600, b) }
func BenchmarkRecursiveLinearSearch_3200(b *testing.B)   { benchmarkRecursiveLinearSearch(3200, b) }
func BenchmarkRecursiveLinearSearch_6400(b *testing.B)   { benchmarkRecursiveLinearSearch(6400, b) }
func BenchmarkRecursiveLinearSearch_12800(b *testing.B)  { benchmarkRecursiveLinearSearch(12800, b) }
func BenchmarkRecursiveLinearSearch_25600(b *testing.B)  { benchmarkRecursiveLinearSearch(25600, b) }
func BenchmarkRecursiveLinearSearch_51200(b *testing.B)  { benchmarkRecursiveLinearSearch(51200, b) }
func BenchmarkRecursiveLinearSearch_102400(b *testing.B) { benchmarkRecursiveLinearSearch(102400, b) }

var loc int
var found bool
