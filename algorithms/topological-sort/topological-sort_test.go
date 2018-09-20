package topologicalsort

import (
	"fmt"
	"reflect"
	"testing"
)

var g1 = Graph{
	[]Node{2},
	[]Node{3},
	[]Node{3, 4},
	[]Node{5},
	[]Node{5},
	[]Node{6, 10},
	[]Node{7},
	[]Node{12},
	[]Node{9},
	[]Node{10},
	[]Node{11},
	[]Node{12},
	[]Node{13},
	[]Node{},
}

// Nodes: 10, Edges: 15
var g10 = Graph{
	[]Node{4, 6, 9},
	[]Node{3, 6, 7, 8},
	[]Node{4, 9},
	[]Node{7, 9},
	[]Node{6},
	[]Node{6, 8},
	[]Node{9},
	[]Node{},
	[]Node{},
	[]Node{},
}

// Nodes: 12, Edges: 22
var g12 = Graph{
	[]Node{4, 6, 9, 10},
	[]Node{3, 6, 7, 8},
	[]Node{4, 9, 10},
	[]Node{7, 9, 10},
	[]Node{6, 10, 11},
	[]Node{6, 8, 11},
	[]Node{9},
	[]Node{},
	[]Node{},
	[]Node{},
	[]Node{},
	[]Node{},
}

// Nodes: 14, Edges: 25
var g14 = Graph{
	[]Node{4, 6, 9, 10},
	[]Node{3, 6, 7, 8, 12},
	[]Node{4, 9, 10},
	[]Node{7, 9, 10},
	[]Node{6, 10, 11},
	[]Node{6, 8, 11},
	[]Node{9},
	[]Node{12},
	[]Node{11},
	[]Node{},
	[]Node{},
	[]Node{},
	[]Node{},
	[]Node{},
}

// Nodes: 16, Edges: 33
var g16 = Graph{
	[]Node{4, 6, 9, 10, 15},
	[]Node{3, 6, 7, 8, 12, 14},
	[]Node{4, 9, 10, 14, 15},
	[]Node{7, 9, 10},
	[]Node{6, 10, 11},
	[]Node{6, 8, 11},
	[]Node{9, 15},
	[]Node{12},
	[]Node{11},
	[]Node{15},
	[]Node{12, 14, 15},
	[]Node{},
	[]Node{},
	[]Node{},
	[]Node{},
	[]Node{},
}

// Nodes: 18, Edges: 43
var g18 = Graph{
	[]Node{4, 6, 9, 10, 15, 16},
	[]Node{3, 6, 7, 8, 12, 14, 16},
	[]Node{4, 9, 10, 14, 15, 16, 17},
	[]Node{7, 9, 10},
	[]Node{6, 10, 11, 17},
	[]Node{6, 8, 11},
	[]Node{9, 15, 17},
	[]Node{12, 17},
	[]Node{11},
	[]Node{15, 16},
	[]Node{12, 14, 15, 17},
	[]Node{16},
	[]Node{},
	[]Node{},
	[]Node{},
	[]Node{},
	[]Node{},
	[]Node{},
}

// Nodes: 20, Edges: 57
var g20 = Graph{
	[]Node{4, 6, 9, 10, 15, 16, 18},
	[]Node{3, 6, 7, 8, 12, 14, 16, 19},
	[]Node{4, 9, 10, 14, 15, 16, 17, 18},
	[]Node{7, 9, 10, 19},
	[]Node{6, 10, 11, 17, 19},
	[]Node{6, 8, 11, 18, 19},
	[]Node{9, 15, 17, 19},
	[]Node{12, 17, 18},
	[]Node{11},
	[]Node{15, 16},
	[]Node{12, 14, 15, 17},
	[]Node{16},
	[]Node{19},
	[]Node{18, 19},
	[]Node{},
	[]Node{},
	[]Node{},
	[]Node{},
	[]Node{},
	[]Node{},
}

var testCases = []struct {
	name   string
	graph  Graph
	sorted []Node
}{
	{
		name:   "simple graph",
		graph:  g1,
		sorted: []Node{0, 1, 8, 2, 9, 3, 4, 5, 6, 10, 7, 11, 12, 13},
	},
}

func TestTopologicalSort(t *testing.T) {
	for _, tc := range testCases {
		sorted := TopologicalSort(tc.graph)
		if got, want := sorted, tc.sorted; reflect.DeepEqual(got, want) == false {
			t.Errorf("\n\t got: %v\n\twant: %v\n", got, want)
		}
	}
}

func benchmarkTopologicalSort(g Graph, b *testing.B) {
	for n := 0; n < b.N; n++ {
		TopologicalSort(g)
	}
}

func BenchmarkTopologicalSort(b *testing.B) {
	graphs := []Graph{g10, g12, g14, g16, g18, g20}
	for i := range graphs {
		b.Run(fmt.Sprintf("_%d", len(graphs[i])), func(b *testing.B) {
			benchmarkTopologicalSort(graphs[i], b)
		})
	}
}
