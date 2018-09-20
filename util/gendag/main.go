package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

// Run generates a Directed Acyclic Graph.
// Algorithm is taken from the accepted answer to this Stack Overflow question:
// https://stackoverflow.com/questions/12790337/generating-a-random-dag
//
// It prints a definition to the console for use in graph algorithms. Additionally
// it creates a graph.dot file for plotting a graphical representation of the
// generated graph.
func run() error {

	const (
		MinPerRank      = 2
		MaxPerRank      = 3
		MinRanks        = 10
		MaxRanks        = 11
		EdgeProbability = 30
	)

	f, err := os.Create("graph.dot")
	if err != nil {
		return err
	}

	ranks := rand.Intn(MaxRanks-MinRanks) + MinRanks
	nodes := 0
	edges := 0

	adj := map[int][]int{}

	fmt.Fprintln(f, "digraph {")

	for i := 0; i < ranks; i++ {
		newNodes := rand.Intn(MaxPerRank-MinPerRank) + MinPerRank

		for j := 0; j < nodes; j++ {
			for k := 0; k < newNodes; k++ {
				if rand.Float64()*100 < EdgeProbability {
					fmt.Fprintf(f, "\t%d  -> %d;\n", j, k+nodes)
					adj[j] = append(adj[j], k+nodes)
					edges++
				}
			}
		}
		nodes += newNodes
	}
	fmt.Fprintln(f, "}")

	fmt.Println()
	fmt.Printf("// Nodes: %d, Edges: %d\n", nodes, edges)
	fmt.Println("var g1 = Graph{")
	for i := 0; i < len(adj); i++ {
		fmt.Print("\t[]Node{")
		for j := range adj[i] {
			fmt.Printf("%d", adj[i][j])
			if j < len(adj[i])-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("},")
	}
	fmt.Println("}")

	err = f.Close()
	return err
}
