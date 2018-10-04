package dagsp

import (
	"math"
)

const (
	inf int = math.MaxInt64
)

// Graph is a directed graph in adjacency-list representation.
// Each entry in the graph represents a node. Its value is a
// list of adjacent nodes.
type Graph [][]Node

// Node is the vertex in a Graph.
type Node struct {
	ID     int
	Weight int
}

// DAGShortestPath takes a directed-acyclic graph and a source node
// and returns an array containing the length of the shortest path
// to each node and an array indicating the nearest neighbour of
// each node.
func DAGShortestPath(g Graph, s int) ([]int, []int) {
	shortest := make([]int, len(g))
	pred := make([]int, len(g))

	for v := range g {
		shortest[v] = inf
		pred[v] = -1
	}
	shortest[s] = 0

	l := TopologicalSort(g)
	for _, u := range l {
		for _, v := range g[u.ID] {
			Relax(u, v, shortest[:], pred[:])
		}
	}
	return shortest, pred
}

// TopologicalSort takes a directed-acyclic graph and returns an
// array of verticies such that `u` appears before `v` in the
// linear order if `(u, v)` is an edge in the graph.
func TopologicalSort(g Graph) []Node {
	var sorted []Node

	inDegree := make([]int, len(g))
	for u := range g {
		for _, v := range g[u] {
			inDegree[v.ID]++
		}
	}

	var next []Node
	for u := range inDegree {
		if inDegree[u] == 0 {
			next = append(next, Node{ID: u})
		}
	}

	for len(next) > 0 {
		u := next[0]
		next = next[1:]

		sorted = append(sorted, u)

		for _, v := range g[u.ID] {
			inDegree[v.ID]--
			if inDegree[v.ID] == 0 {
				next = append(next, v)
			}

		}
	}

	return sorted
}

// Relax takes two ajacent nodes and the current shortest path and
// route between them and determines whether the new shortes path
// is less than the current. If so it modifies the slices containing
// the shortest path weighting, s, and route, p.
func Relax(u, v Node, s, p []int) {
	if s[u.ID]+v.Weight < s[v.ID] && s[u.ID]+v.Weight >= 0 {
		s[v.ID] = s[u.ID] + v.Weight
		p[v.ID] = u.ID
	}
}
