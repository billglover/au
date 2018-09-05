package binarysearch

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

func TestBinarySearch(t *testing.T) {
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

func benchmarkBinarySearch(i int, b *testing.B) {
	var l int
	var f bool

	rand.Seed(time.Now().UnixNano())
	A := rand.Perm(i)
	x := A[rand.Intn(i)]

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		l, f = BinarySearch(A, x)
	}
	loc, found = l, f
}

func BenchmarkBinarySearch_100(b *testing.B)    { benchmarkBinarySearch(100, b) }
func BenchmarkBinarySearch_200(b *testing.B)    { benchmarkBinarySearch(200, b) }
func BenchmarkBinarySearch_400(b *testing.B)    { benchmarkBinarySearch(400, b) }
func BenchmarkBinarySearch_800(b *testing.B)    { benchmarkBinarySearch(800, b) }
func BenchmarkBinarySearch_1600(b *testing.B)   { benchmarkBinarySearch(1600, b) }
func BenchmarkBinarySearch_3200(b *testing.B)   { benchmarkBinarySearch(3200, b) }
func BenchmarkBinarySearch_6400(b *testing.B)   { benchmarkBinarySearch(6400, b) }
func BenchmarkBinarySearch_12800(b *testing.B)  { benchmarkBinarySearch(12800, b) }
func BenchmarkBinarySearch_25600(b *testing.B)  { benchmarkBinarySearch(25600, b) }
func BenchmarkBinarySearch_51200(b *testing.B)  { benchmarkBinarySearch(51200, b) }
func BenchmarkBinarySearch_102400(b *testing.B) { benchmarkBinarySearch(102400, b) }

var loc int
var found bool
