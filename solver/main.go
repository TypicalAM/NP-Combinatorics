package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Mandala/go-log"
	"github.com/TypicalAM/NP-Combinatorics/src/input"
	"github.com/TypicalAM/NP-Combinatorics/src/solutions"
)

// Driver code
func main() {
	logger := log.New(os.Stdout)

	destinationFlag := flag.String("load", "data/example.json", "Select the destination of the source file of the graph")
	flag.Parse()

	data, err := input.LoadData(*destinationFlag)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("The distance matrix for the input graph")

	for row := 0; row < len(data.Distances); row++ {
		for column := 0; column < len(data.Distances); column++ {
			fmt.Print(data.Distances[row][column], " ")
		}
		fmt.Println("")
	}

	_ = solutions.BruteForce(logger, data)
	_ = solutions.Greedy(logger, data)
}
