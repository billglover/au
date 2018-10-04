package dagsp

import (
	"reflect"
	"testing"
)

var g1 = Graph{
	[]Node{{ID: 2, Weight: 1}},
	[]Node{{ID: 3, Weight: 1}},
	[]Node{{ID: 3, Weight: 1}, {ID: 4, Weight: 1}},
	[]Node{{ID: 5, Weight: 1}},
	[]Node{{ID: 5, Weight: 1}},
	[]Node{{ID: 6, Weight: 1}, {ID: 10, Weight: 1}},
	[]Node{{ID: 7, Weight: 1}},
	[]Node{{ID: 12, Weight: 1}},
	[]Node{{ID: 9, Weight: 1}},
	[]Node{{ID: 10, Weight: 1}},
	[]Node{{ID: 11, Weight: 1}},
	[]Node{{ID: 12, Weight: 1}},
	[]Node{{ID: 13, Weight: 1}},
	[]Node{},
}

var testCases = []struct {
	name     string
	graph    Graph
	sorted   []Node
	shortest []int
	pred     []int
}{
	{
		name:     "simple graph",
		graph:    g1,
		sorted:   []Node{{ID: 0}, {ID: 1}, {ID: 8}, {ID: 2, Weight: 1}, {ID: 9, Weight: 1}, {ID: 3, Weight: 1}, {ID: 4, Weight: 1}, {ID: 5, Weight: 1}, {ID: 6, Weight: 1}, {ID: 10, Weight: 1}, {ID: 7, Weight: 1}, {ID: 11, Weight: 1}, {ID: 12, Weight: 1}, {ID: 13, Weight: 1}},
		shortest: []int{0, inf, 1, 2, 2, 3, 4, 5, inf, inf, 4, 5, 6, 7},
		pred:     []int{-1, -1, 0, 2, 2, 3, 5, 6, -1, -1, 5, 10, 7, 12},
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

func TestDAGShortestPath(t *testing.T) {
	for _, tc := range testCases {
		shortest, pred := DAGShortestPath(tc.graph, 0)
		if got, want := shortest, tc.shortest; reflect.DeepEqual(got, want) == false {
			t.Errorf("unexpected shortest path weight\n\t got: %v\n\twant: %v\n", got, want)
		}

		if got, want := pred, tc.pred; reflect.DeepEqual(got, want) == false {
			t.Errorf("unexpected shortest path route\n\t got: %v\n\twant: %v\n", got, want)
		}
	}
}

// // Nodes: 10, Edges: 15
// var g10 = Graph{
// 	[]Node{{ID: 4}, {ID: 6}, {ID: 9}},
// 	[]Node{{ID: 3}, {ID: 6}, {ID: 7}, {ID: 8}},
// 	[]Node{{ID: 4}, {ID: 9}},
// 	[]Node{{ID: 7}, {ID: 9}},
// 	[]Node{{ID: 6}},
// 	[]Node{{ID: 6}, {ID: 8}},
// 	[]Node{{ID: 9}},
// 	[]Node{},
// 	[]Node{},
// 	[]Node{},
// }

// // Nodes: 12, Edges: 22
// var g12 = Graph{
// 	[]Node{{ID: 4}, {ID: 6}, {ID: 9}, {ID: 10}},
// 	[]Node{{ID: 3}, {ID: 6}, {ID: 7}, {ID: 8}},
// 	[]Node{{ID: 4}, {ID: 9}, {ID: 10}},
// 	[]Node{{ID: 7}, {ID: 9}, {ID: 10}},
// 	[]Node{{ID: 6}, {ID: 10}, {ID: 11}},
// 	[]Node{{ID: 6}, {ID: 8}, {ID: 11}},
// 	[]Node{{ID: 9}},
// 	[]Node{},
// 	[]Node{},
// 	[]Node{},
// 	[]Node{},
// 	[]Node{},
// }

// // Nodes: 14, Edges: 25
// var g14 = Graph{
// 	[]Node{{ID: 4}, {ID: 6}, {ID: 9}, {ID: 10}},
// 	[]Node{{ID: 3}, {ID: 6}, {ID: 7}, {ID: 8}, {ID: 12}},
// 	[]Node{{ID: 4}, {ID: 9}, {ID: 10}},
// 	[]Node{{ID: 7}, {ID: 9}, {ID: 10}},
// 	[]Node{{ID: 6}, {ID: 10}, {ID: 11}},
// 	[]Node{{ID: 6}, {ID: 8}, {ID: 11}},
// 	[]Node{{ID: 9}},
// 	[]Node{{ID: 12}},
// 	[]Node{{ID: 11}},
// 	[]Node{},
// 	[]Node{},
// 	[]Node{},
// 	[]Node{},
// 	[]Node{},
// }

// // Nodes: 16, Edges: 33
// var g16 = Graph{
// 	[]Node{{ID: 4}, {ID: 6}, {ID: 9}, {ID: 10}, {ID: 15}},
// 	[]Node{{ID: 3}, {ID: 6}, {ID: 7}, {ID: 8}, {ID: 12}, {ID: 14}},
// 	[]Node{{ID: 4}, {ID: 9}, {ID: 10}, {ID: 14}, {ID: 15}},
// 	[]Node{{ID: 7}, {ID: 9}, {ID: 10}},
// 	[]Node{{ID: 6}, {ID: 10}, {ID: 11}},
// 	[]Node{{ID: 6}, {ID: 8}, {ID: 11}},
// 	[]Node{{ID: 9}, {ID: 15}},
// 	[]Node{{ID: 12}},
// 	[]Node{{ID: 11}},
// 	[]Node{{ID: 15}},
// 	[]Node{{ID: 12}, {ID: 14}, {ID: 15}},
// 	[]Node{},
// 	[]Node{},
// 	[]Node{},
// 	[]Node{},
// 	[]Node{},
// }

// // Nodes: 18, Edges: 43
// var g18 = Graph{
// 	[]Node{{ID: 4}, {ID: 6}, {ID: 9}, {ID: 10}, {ID: 15}, {ID: 16}},
// 	[]Node{{ID: 3}, {ID: 6}, {ID: 7}, {ID: 8}, {ID: 12}, {ID: 14}, {ID: 16}},
// 	[]Node{{ID: 4}, {ID: 9}, {ID: 10}, {ID: 14}, {ID: 15}, {ID: 16}, {ID: 17}},
// 	[]Node{{ID: 7}, {ID: 9}, {ID: 10}},
// 	[]Node{{ID: 6}, {ID: 10}, {ID: 11}, {ID: 17}},
// 	[]Node{{ID: 6}, {ID: 8}, {ID: 11}},
// 	[]Node{{ID: 9}, {ID: 15}, {ID: 17}},
// 	[]Node{{ID: 12}, {ID: 17}},
// 	[]Node{{ID: 11}},
// 	[]Node{{ID: 15}, {ID: 16}},
// 	[]Node{{ID: 12}, {ID: 14}, {ID: 15}, {ID: 17}},
// 	[]Node{{ID: 16}},
// 	[]Node{},
// 	[]Node{},
// 	[]Node{},
// 	[]Node{},
// 	[]Node{},
// 	[]Node{},
// }

// // Nodes: 20, Edges: 57
// var g20 = Graph{
// 	[]Node{{ID: 4}, {ID: 6}, {ID: 9}, {ID: 10}, {ID: 15}, {ID: 16}, {ID: 18}},
// 	[]Node{{ID: 3}, {ID: 6}, {ID: 7}, {ID: 8}, {ID: 12}, {ID: 14}, {ID: 16}, {ID: 19}},
// 	[]Node{{ID: 4}, {ID: 9}, {ID: 10}, {ID: 14}, {ID: 15}, {ID: 16}, {ID: 17}, {ID: 18}},
// 	[]Node{{ID: 7}, {ID: 9}, {ID: 10}, {ID: 19}},
// 	[]Node{{ID: 6}, {ID: 10}, {ID: 11}, {ID: 17}, {ID: 19}},
// 	[]Node{{ID: 6}, {ID: 8}, {ID: 11}, {ID: 18}, {ID: 19}},
// 	[]Node{{ID: 9}, {ID: 15}, {ID: 17}, {ID: 19}},
// 	[]Node{{ID: 12}, {ID: 17}, {ID: 18}},
// 	[]Node{{ID: 11}},
// 	[]Node{{ID: 15}, {ID: 16}},
// 	[]Node{{ID: 12}, {ID: 14}, {ID: 15}, {ID: 17}},
// 	[]Node{{ID: 16}},
// 	[]Node{{ID: 19}},
// 	[]Node{{ID: 18}, {ID: 19}},
// 	[]Node{},
// 	[]Node{},
// 	[]Node{},
// 	[]Node{},
// 	[]Node{},
// 	[]Node{},
// }
