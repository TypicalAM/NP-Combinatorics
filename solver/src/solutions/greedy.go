package solutions

import (
	"math"
	"time"

	"github.com/Mandala/go-log"
	g "github.com/TypicalAM/NP-Combinatorics/src/graph"
	"github.com/TypicalAM/NP-Combinatorics/src/timetrack"
)

// Solve the problem using the greedy method
func greedy(logger *log.Logger, graph g.Graph) int {
	defer timetrack.TimeTrack(logger, time.Now(), "Greedy solution")
	logger.Info("Running the greedy solution")

	var total, minIndex, vertex int

	vertices := make([]int, len(graph.Distances))
	visited := make([]bool, len(vertices))

	for i := range graph.Distances {
		vertices[i] = i
	}

	for i := 0; i < len(vertices)-1; i++ {
		min := math.MaxInt
		visited[vertex] = true

		for ind, elem := range graph.Distances[vertex] {
			if elem != 0 && elem < min && !visited[ind] {
				min = elem
				minIndex = ind
			}
		}

		vertex = minIndex
		total += min
	}

	total += graph.Distances[len(graph.Distances)-1][0]

	return total
}
