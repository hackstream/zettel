package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hackstream/zettel/internal/pipeline"
	"github.com/yourbasic/graph"
)

// Node represents each node in the graph
type Node struct {
	ID    string  `json:"id"`
	Label string  `json:"label"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Size  int     `json:"size"`
}

// Edge represents a connection between two nodes in the graph
type Edge struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
}

// GraphData is representation of the graph in JSON
type GraphData struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

// MakeGraphData returns a GraphData from the given posts and graph
func MakeGraphData(posts []pipeline.Post, graph *graph.Mutable) GraphData {
	graphData := GraphData{
		Nodes: make([]Node, 0, len(posts)),
		Edges: make([]Edge, 0),
	}

	rand.Seed(time.Now().UnixNano())

	// Loop over the graph and add nodes & edges
	for i := 0; i < len(posts); i++ {
		p := posts[i]

		node := Node{
			ID:    fmt.Sprintf("nodeid-%d", i),
			Label: p.Meta.Title,
			Size:  5,
			X:     rand.Float64(),
			Y:     rand.Float64(),
		}

		graphData.Nodes = append(graphData.Nodes, node)

		for j := i + 1; j < len(posts); j++ {
			if graph.Edge(i, j) {
				edge := Edge{
					ID:     fmt.Sprintf("edge-%d-%d", i, j),
					Source: fmt.Sprintf("nodeid-%d", i),
					Target: fmt.Sprintf("nodeid-%d", j),
				}
				graphData.Edges = append(graphData.Edges, edge)
			}
		}
	}

	return graphData
}
