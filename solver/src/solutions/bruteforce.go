package solutions

import (
	"math"
	"time"

	"github.com/Mandala/go-log"
	g "github.com/TypicalAM/NP-Combinatorics/src/graph"
	"github.com/TypicalAM/NP-Combinatorics/src/timetrack"
	"gonum.org/v1/gonum/stat/combin"
)

// Solve the problem using the bruteforce method
func bruteforce(logger *log.Logger, graph g.Graph) int {
	defer timetrack.TimeTrack(logger, time.Now(), "Bruteforce solution")
	logger.Info("---- Running the bruteforce solution ----")

	logger.Infof("Generating permutations")

	// Create the permutation generator instance
	vertices := len(graph.Distances)
	permGenerator := combin.NewPermutationGenerator(vertices, vertices)
	permutation := make([]int, vertices)
	min := math.MaxInt

	logger.Infof("Generated the permutation object")

	// Find the minimum distance permutation
	for permGenerator.Next() {
		permGenerator.Permutation(permutation)
		var total int
		// Add the distance between the first element and the last element of the permutation
		total += graph.Distances[permutation[0]][permutation[len(graph.Distances)-1]]

		// Add subsequent elements of the slice
		for ind := 0; ind < len(permutation)-1; ind++ {
			total += graph.Distances[permutation[ind]][permutation[ind+1]]
		}

		// Make it the new minimum if it is the best distance
		if total < min {
			min = total
		}
	}

	return min
}
