package main

import (
	"os"

	"github.com/Mandala/go-log"
	"github.com/TypicalAM/NP-Combinatorics/src/input"
	"github.com/TypicalAM/NP-Combinatorics/src/settings"
	"github.com/TypicalAM/NP-Combinatorics/src/solutions"
)

// Driver code
func main() {
	logger := log.New(os.Stdout)

	appSettings, err := settings.GetSettingsFromCmdArgs()
	if err != nil {
		logger.Fatal(err)
	}

	data, err := input.LoadData(appSettings.Destination)
	if err != nil {
		logger.Fatal(err)
	}

	solutions.RunSolutions(logger, appSettings, data)
}
