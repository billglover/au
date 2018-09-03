# Bad Factorial

## Pseudocode

**Procedure:** `BadFactorial(n)`

**Input:** An integer `n >= 0`.

**Output:** The value of `n!`.

**Implementation:**

1. If n=0, then return 1 as the output.
2. Otherwise, return `BadFactorial(n+1) / (n+1)`.

## Results

```plain
$ go test -v
=== RUN   TestBadFactorial
=== RUN   TestBadFactorial/0!
=== RUN   TestBadFactorial/1!
runtime: goroutine stack exceeds 1000000000-byte limit
fatal error: stack overflow
```

The Bad Factorial algorithm results in an infinite loop because the recursive function never converges.
