package solutions

import (
	"fmt"

	"github.com/Mandala/go-log"
	g "github.com/TypicalAM/NP-Combinatorics/src/graph"
	"github.com/TypicalAM/NP-Combinatorics/src/settings"
)

// Display the graph in a grid
func showGraph(logger *log.Logger, data g.Graph) {
	logger.Info("The distance matrix for the input graph")

	for row := 0; row < len(data.Distances); row++ {
		for column := 0; column < len(data.Distances); column++ {
			fmt.Print(data.Distances[row][column], " ")
		}
		fmt.Println()
	}
}

// Run the selected solutions
func RunSolutions(logger *log.Logger, appSettings settings.Settings, data g.Graph) {
	var greedyResult int

	if appSettings.ShowGraph {
		showGraph(logger, data)
	}

	if appSettings.UseGreedy {
		if appSettings.Parralel {
			greedyResult = greedyParralel(logger, data)
			logger.Info("Best path using parallel greedy:", greedyResult)
		} else {
			greedyResult = greedySequential(logger, data)
			logger.Warn("The backtracking data may not be as accurate when not using the parallel option")
			logger.Info("Best path using sequential greedy:", greedyResult)
		}
	}

	if appSettings.UseBruteforce {
		logger.Info("Best path using bruteforce:", bruteforce(logger, data))
	}

	if appSettings.UseBacktracking {
		logger.Info("Best path using backtracking:", backtracking(logger, data, greedyResult))
	}

}
