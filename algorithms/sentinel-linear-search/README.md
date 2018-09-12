# Sentinel Linear Search

## Pseudocode

**Procedure:** `SentinelLinearSearch(A, n, x)`

**Inputs:**

* `A`: an array
* `n`: the number of elements in A to search through
* `x`: the value being searched for

**Output:** Either an index `i` for which `A[i] = x`, or the special value `NOT-FOUND`, which could be any invalid index into the array, such as `0` or any negative integer.

**Implementation:**

1. Save `A[n]` into `last` and then put `x` into `A[n]`.
2. Set `i` to `1`.
3. While `A[i] != x`, do the following:
   1. Increment `i`.
4. Restore `A[n]` from `last`.
5. If `i < n` or `A[n] == x`, the return the value of `i` as the output.
6. Otherwise, return `NOT-FOUND` as the output.

## Implementation

* Instead of a magic number, we take advantage of Go's multiple return types and use a boolean to indicate whether the result was found.
* Arrays are zero indexed in Go so we adjust the loop accordingly.
* We infer `n` based on the length of `A`.

## Results

**Benchmark:**

```plain
$ go test -run=X -bench=Linear -benchmem
goos: darwin
goarch: amd64
pkg: github.com/billglover/au/algorithms/sentinel-linear-search
BenchmarkLinearSearch_100-4             30000000                78.3 ns/op             0 B/op          0 allocs/op
BenchmarkLinearSearch_200-4             10000000               113 ns/op               0 B/op          0 allocs/op
BenchmarkLinearSearch_400-4             10000000               195 ns/op               0 B/op          0 allocs/op
BenchmarkLinearSearch_800-4              5000000               320 ns/op               0 B/op          0 allocs/op
BenchmarkLinearSearch_1600-4             2000000               628 ns/op               0 B/op          0 allocs/op
BenchmarkLinearSearch_3200-4             1000000              1241 ns/op               0 B/op          0 allocs/op
BenchmarkLinearSearch_6400-4              500000              2666 ns/op               0 B/op          0 allocs/op
BenchmarkLinearSearch_12800-4             300000              5090 ns/op               0 B/op          0 allocs/op
BenchmarkLinearSearch_25600-4             100000             11730 ns/op               0 B/op          0 allocs/op
BenchmarkLinearSearch_51200-4              50000             24627 ns/op               0 B/op          0 allocs/op
BenchmarkLinearSearch_102400-4             30000             52495 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/billglover/au/algorithms/sentinel-linear-search      17.607s
```

**Complexity:**

| Running Time | Complexity |
|--------------|------------|
| Worst        | Θ(n)       |
| Best         | Θ(1)       |
| Typical      | Θ(n)       |

Worst case performance of the algorithm requires searching the full array. In this case we are searching for a value we know does not exist.

![Time Complexity: Linear Search](img/complexity_time_worst.png)

For comparison the best case performance of the algorithm doesn't vary with the array size. This is because the algorithm returns as soon as the value is found.

![Time Complexity: Linear Search](img/complexity_time_best.png)

With typical run-times, we see variable performance as it is rare that we end up searching the full array for the value in question. Unlike [Linear Search](algorithms/linear-search/), we return as soon as we have found the value we are looking for.

![Time Complexity: Linear Search](img/complexity_time_typical.png)
