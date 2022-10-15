package input

import (
	g "github.com/TypicalAM/NP-Combinatorics/src/graph"
)

type NodeData struct {
	ID int `json:"id"`
}

type EdgesData struct {
	ID     int `json:"id"`
	Weight int `json:"weight"`
}

type GraphData struct {
	Adjacency  [][]EdgesData `json:"adjacency"`
	Graph      any           `json:"graph"`
	Nodes      []NodeData    `json:"nodes"`
	Directed   bool          `json:"directed"`
	Multigraph bool          `json:"multigraph"`
}

// Convert the GraphData struct into a proper representation
func (gdata GraphData) convertToRepresentation() g.Graph {
	matrix := make([][]int, len(gdata.Nodes))
	for i := range matrix {
		matrix[i] = make([]int, len(gdata.Nodes))
	}

	for from, vertexAdj := range gdata.Adjacency {
		for _, edge := range vertexAdj {
			matrix[from][edge.ID] = edge.Weight
		}
	}

	return g.Graph{Distances: matrix}
}
