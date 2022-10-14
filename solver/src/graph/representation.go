package graph

// Represent the graph in a adjecency matrix
type Graph struct {
	Edges     [][]bool `json:"edges"`
	Distances [][]int  `json:"distances"`
}
