# DAG Shortest Paths

## Pseudocode

**Procedure:** `DAGShortestPath(G, s)`

**Inputs:**

* `G`: a weighted directed acyclic graph contain a set `V` of `n` vertices and a set `E` of `m` directed edges.
* `s`: a source vertex in `V`.

**Result:** For each non-source vertex `v` in `V`, `shortest[v]` is the weight `sp(s, v)` of a shortest path from `s` to `v` and `pred[v]` is the vertex preceding `v` on some shortest path. For the source vertex `s`, `shortest[s] = 0` and `pred[s] = nil`. I there is no path from `s` to `v`, then `shortest[v] = ∞` and `pred[v] = nil`.

**Implementation:**

1. Call `TopologicalSort(G)` and set `l` to be the linear order of the vertices returned by the call.
2. Set `shortest[v]` to `∞` for each vertex `v` except `s`, set `shortest[s]` to `0`, and set `pred[v]` to `nil` for each vertex `v`.
3. For each vertex `u`, taken in the order given by `l`:
   1. For each vertex `v` adjacent to `u`:
      1. Call `Relax(u, v)`.

**Procedure:** `Relax(u, v)`

**Inputs:** `u`, `v`: vertices such that there is an edge `(u, v)`.

**Result:** The value of `shortest[v]` might decrease, and if it does, then `pred[v]` becomes `u`.

**Implementation:**

1. If `shortest[u] + weight(u, v) < shortest[v]`, then set `shortest[v]` to `shortest[u] + weight(u, v)` and set `pred[v]` to `u`.

## Implementation Notes

* We define `inf` as the max int value.
* During the Relax procedure we have to check we don't wrap around beyond `inf`. We do this by ensuring the value remains positive.
