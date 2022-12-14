package settings

import (
	"flag"
	"fmt"
	"strings"
)

// A struct to hold the app settings
type Settings struct {
	Destination     string
	Parralel        bool
	ShowGraph       bool
	UseBruteforce   bool
	UseGreedy       bool
	UseBacktracking bool
}

// Check if flags are in the correct format
func GetSettingsFromCmdArgs() (Settings, error) {
	solvPtr := flag.String("solvers", "bruteforce,backtracking,greedy", "Choose the solvers for the presented problem, delimited by commas")
	destPtr := flag.String("load", "data/example.json", "Select the destination of the source file of the graph")
	showPtr := flag.Bool("show", false, "Show the imported graph before computing solutions")
	parrPtr := flag.Bool("parralel", true, "Use parralel processing for the greedy solution")

	flag.Parse()

	settings := Settings{}
	settings.Destination = *destPtr
	settings.ShowGraph = *showPtr
	settings.Parralel = *parrPtr

	for _, solver := range strings.Split(*solvPtr, ",") {
		switch solver {
		case "bruteforce":
			settings.UseBruteforce = true
		case "backtracking":
			settings.UseBacktracking = true
		case "greedy":
			settings.UseGreedy = true
		default:
			return Settings{}, fmt.Errorf("unrecognised argument to solvers: %s", solver)
		}
	}

	return settings, nil
}
