package badfactorial

// BadFactorial demonstrates how not to do recursion as it results in an
// infinite loop.
func BadFactorial(n int) int {
	if n == 0 {
		return 1
	}
	return BadFactorial(n+1) / (n + 1)
}
