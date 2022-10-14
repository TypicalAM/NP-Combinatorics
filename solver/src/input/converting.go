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
func (gdata GraphData) convertToRepresentation() g.Graph { return g.Graph{} }
