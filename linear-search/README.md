# Linear Search

## Pseudocode

**Procedure:** `LinearSearch(A, n, x)`

**Inputs:**

* `A`: an array
* `n`: the number of elements in A to search through
* `x`: the value being searched for

**Output:** Either an index `i` for which `A[i] = x`, or the special value `NOT-FOUND`, which could be any invalid index into the array, such as `0` or any negative integer.

**Implementation:**

1. Set `answer` to `NOT-FOUND`.
2. For each index `i`, going from `1` to `n`, in order:
   1. If `A[i] = x`, then set answer to the value of `i`.
3. Return the value of answer as the output.

## Implementation

* Instead of a magic number, we take advantage of Go's multiple return types and use a boolean to indicate whether the result was found.
* Arrays are zero indexed in Go so we adjust the loop accordingly.
* We infer `n` based on the length of `A`.

## Results
