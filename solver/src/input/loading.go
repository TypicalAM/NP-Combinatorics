package input

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	g "github.com/TypicalAM/NP-Combinatorics/src/graph"
)

// Load the data from a file (filepath) and convert it to a proper graph
func LoadData(filepath string) (g.Graph, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Print(err)
		return g.Graph{}, errors.New("couldn't load the file")
	}

	data := GraphData{}
	if err = json.Unmarshal(file, &data); err != nil {
		return g.Graph{}, errors.New("couldn't convert the file from json")
	}

	return data.convertToRepresentation(), nil
}
