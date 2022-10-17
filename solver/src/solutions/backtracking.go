package solutions

import (
	"math"
	"time"

	"github.com/Mandala/go-log"
	g "github.com/TypicalAM/NP-Combinatorics/src/graph"
	"github.com/TypicalAM/NP-Combinatorics/src/timetrack"
)

// Returns the minimal value of the two integers provided
func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// A recursive helper function for the backtracking method
func helper(graph g.Graph, visited *[]bool, currentPosition, size, count, dist int, ans *int) {
	// If the path is full (no new distinct values can be added) check if it is the smallest answer
	if count == size && graph.Distances[currentPosition][0] != 0 {
		*ans = min(*ans, dist+graph.Distances[currentPosition][0])
		return
	}

	visitedValue := *visited

	// Backtracking
	for i := 0; i < size; i++ {
		if !visitedValue[i] && graph.Distances[currentPosition][i] != 0 {
			visitedValue[i] = true

			helper(graph, visited, i, size, count+1, dist+graph.Distances[currentPosition][i], ans)

			visitedValue[i] = false
		}
	}
}

// Solve the problem using the backtracking method
func backtracking(logger *log.Logger, graph g.Graph) int {
	defer timetrack.TimeTrack(logger, time.Now(), "Backtracking solution")
	logger.Info("---- Running the backtracking solution ----")

	// Visited[i] denotes if the i-th vertex has been visited
	visited := make([]bool, len(graph.Distances))
	visited[0] = true

	// Initialize the result and pass it onto the helper function to mutate it
	result := math.MaxInt
	helper(graph, &visited, 0, len(graph.Distances), 1, 0, &result)

	return result
}
