package factorial

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	n         int
	factorial int
}{
	{n: 0, factorial: 1},
	{n: 1, factorial: 1},
	{n: 2, factorial: 2},
	{n: 3, factorial: 6},
	{n: 20, factorial: 2432902008176640000},
}

func TestFactorial(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d!", tc.n), func(t *testing.T) {
			factorial := Factorial(tc.n)

			if got, want := factorial, tc.factorial; got != want {
				t.Errorf("unexpected result: got %d, want %d", got, want)
			}
		})
	}

}

func benchmarkFactorial(i int, b *testing.B) {
	var f int

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		f = Factorial(i)
	}
	factorial = f
}

func BenchmarkFactorial_100(b *testing.B)    { benchmarkFactorial(100, b) }
func BenchmarkFactorial_200(b *testing.B)    { benchmarkFactorial(200, b) }
func BenchmarkFactorial_400(b *testing.B)    { benchmarkFactorial(400, b) }
func BenchmarkFactorial_800(b *testing.B)    { benchmarkFactorial(800, b) }
func BenchmarkFactorial_1600(b *testing.B)   { benchmarkFactorial(1600, b) }
func BenchmarkFactorial_3200(b *testing.B)   { benchmarkFactorial(3200, b) }
func BenchmarkFactorial_6400(b *testing.B)   { benchmarkFactorial(6400, b) }
func BenchmarkFactorial_12800(b *testing.B)  { benchmarkFactorial(12800, b) }
func BenchmarkFactorial_25600(b *testing.B)  { benchmarkFactorial(25600, b) }
func BenchmarkFactorial_51200(b *testing.B)  { benchmarkFactorial(51200, b) }
func BenchmarkFactorial_102400(b *testing.B) { benchmarkFactorial(102400, b) }

var factorial int
