package badfactorial

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

func TestBadFactorial(t *testing.T) {
	// Remove the following statement to run the tests. They have been disabled
	// because they result in an infinite loop.
	t.SkipNow()

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d!", tc.n), func(t *testing.T) {
			factorial := BadFactorial(tc.n)

			if got, want := factorial, tc.factorial; got != want {
				t.Errorf("unexpected result: got %d, want %d", got, want)
			}
		})
	}

}
