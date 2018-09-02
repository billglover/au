package betterlinearsearch

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
			loc, found := BetterLinearSearch(tc.A, tc.x)

			if got, want := found, tc.found; got != want {
				t.Errorf("unexpected found: got %v, want %v", got, want)
			}

			if got, want := loc, tc.i; got != want {
				t.Errorf("unexpected index: got %d, want %d", got, want)
			}
		})
	}

}

func benchmarkLinearSearch(i int, b *testing.B) {
	var l int
	var f bool

	rand.Seed(time.Now().UnixNano())
	A := rand.Perm(i)
	x := -1

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		l, f = BetterLinearSearch(A, x)
	}
	loc, found = l, f
}

func BenchmarkLinearSearch_100(b *testing.B)    { benchmarkLinearSearch(100, b) }
func BenchmarkLinearSearch_200(b *testing.B)    { benchmarkLinearSearch(200, b) }
func BenchmarkLinearSearch_400(b *testing.B)    { benchmarkLinearSearch(400, b) }
func BenchmarkLinearSearch_800(b *testing.B)    { benchmarkLinearSearch(800, b) }
func BenchmarkLinearSearch_1600(b *testing.B)   { benchmarkLinearSearch(1600, b) }
func BenchmarkLinearSearch_3200(b *testing.B)   { benchmarkLinearSearch(3200, b) }
func BenchmarkLinearSearch_6400(b *testing.B)   { benchmarkLinearSearch(6400, b) }
func BenchmarkLinearSearch_12800(b *testing.B)  { benchmarkLinearSearch(12800, b) }
func BenchmarkLinearSearch_25600(b *testing.B)  { benchmarkLinearSearch(25600, b) }
func BenchmarkLinearSearch_51200(b *testing.B)  { benchmarkLinearSearch(51200, b) }
func BenchmarkLinearSearch_102400(b *testing.B) { benchmarkLinearSearch(102400, b) }

var loc int
var found bool
