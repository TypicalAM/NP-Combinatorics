package solutions

import (
	"math"
	"time"

	"github.com/Mandala/go-log"
	g "github.com/TypicalAM/NP-Combinatorics/src/graph"
	"github.com/TypicalAM/NP-Combinatorics/src/timetrack"
)

// Generate permutations of a slice of ints, returns all possible combinations of the elements in the slice
func permute(arr []int) [][]int {
	var helper func([]int, int)

	result := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			result = append(result, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				arr[0], arr[n-1] = arr[n-1], arr[0]
			}
		}
	}
	helper(arr, len(arr))

	return result
}

// Solve the problem using the bruteforce method
func bruteforce(logger *log.Logger, graph g.Graph) int {
	defer timetrack.TimeTrack(logger, time.Now(), "Bruteforce solution")
	logger.Info("Running the bruteforce solution")

	// Get a slice of all the vertices
	vertices := make([]int, len(graph.Distances))
	for i := range graph.Distances {
		vertices[i] = i
	}

	logger.Infof("Generating permutations")

	// Generate all the possible permutations of the vertex set
	permutations := permute(vertices[1:])
	min := math.MaxInt
	minIndex := 0

	logger.Infof("Generated %d permutations", len(graph.Distances)*len(permutations))

	// Find the minimum distance permutation
	for permIndex, permutation := range permutations {
		var total int
		// Add the distance between 0 and the last element
		total += graph.Distances[0][permutation[len(graph.Distances)-2]]

		// Add the distance between 0 and the first element
		total += graph.Distances[0][permutation[0]]

		// Add subsequent elements of the slice
		for ind := 0; ind < len(permutation)-1; ind++ {
			total += graph.Distances[permutation[ind]][permutation[ind+1]]
		}

		// Make it the new minimum if it is the best distance
		if total < min {
			min = total
			minIndex = permIndex
		}
	}

	logger.Infof("The best route is %v with 0 as the starting city", permutations[minIndex])

	return min
}
