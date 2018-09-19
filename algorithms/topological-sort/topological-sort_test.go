package topologicalsort

import (
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
