package topologicalsort

// Graph is a directed graph in adjacency-list representation.
// Each entry in the graph represents a node. Its value is a
// list of adjacent nodes.
type Graph [][]Node

// Node is the vertex in a Graph.
type Node int

// TopologicalSort takes a directed-acyclic graph and returns an
// array of verticies such that `u` appears before `v` in the
// linear order if `(u, v)` is an edge in the graph.
func TopologicalSort(g Graph) []Node {
	var sorted []Node

	inDegree := make([]int, len(g))
	for u := range g {
		for _, v := range g[u] {
			inDegree[v]++
		}
	}

	var next []Node
	for u := range inDegree {
		if inDegree[u] == 0 {
			next = append(next, Node(u))
		}
	}

	for len(next) > 0 {
		u := next[0]
		next = next[1:]

		sorted = append(sorted, u)

		for _, v := range g[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				next = append(next, v)
			}

		}
	}

	return sorted
}
