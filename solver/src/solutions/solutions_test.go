package solutions

import (
	"fmt"
	"os"
	"testing"

	"github.com/Mandala/go-log"
	g "github.com/TypicalAM/NP-Combinatorics/src/graph"
	"github.com/TypicalAM/NP-Combinatorics/src/input"
)

type testinfo struct {
	logger    *log.Logger
	testcases []testcase
}

type testcase struct {
	graph    g.Graph
	vertices int
	want     int
}

// Generate the testcases
func genTestcases(isSmall bool) testinfo {
	var verticesNum, expectedResult []int
	var testFiles []string
	if isSmall {
		verticesNum = []int{8, 9, 10, 11}
		expectedResult = []int{455, 425, 412, 366}
		testFiles = []string{
			"../../data/test/8_test.json",
			"../../data/test/9_test.json",
			"../../data/test/10_test.json",
			"../../data/test/11_test.json",
		}
	} else {
		verticesNum = []int{25, 50, 75, 100}
		expectedResult = []int{759, 1255, 1827, 2302}
		testFiles = []string{
			"../../data/test/25_test.json",
			"../../data/test/50_test.json",
			"../../data/test/75_test.json",
			"../../data/test/100_test.json",
		}
	}
	tester := testinfo{}
	tester.logger = log.New(os.Stdout).Quiet()

	tester.testcases = make([]testcase, len(testFiles))

	for ind, path := range testFiles {
		graph, err := input.LoadData(path)
		if err != nil {
			panic("can't load the appropriate test file")
		}

		tester.testcases[ind] = testcase{
			graph:    graph,
			vertices: verticesNum[ind],
			want:     expectedResult[ind],
		}
	}

	return tester
}

// Test the bruteforce solution
func TestBruteforce(t *testing.T) {
	tester := genTestcases(true)

	for _, tt := range tester.testcases {
		testname := fmt.Sprintf("%dv,%d", tt.vertices, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := bruteforce(tester.logger, tt.graph)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Test the backtracking solution
func TestBacktracking(t *testing.T) {
	tester := genTestcases(true)

	for _, tt := range tester.testcases {
		testname := fmt.Sprintf("%dv,%d", tt.vertices, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := backtracking(tester.logger, tt.graph)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Test the greedy solution
func TestGreedy(t *testing.T) {
	tester := genTestcases(false)

	for _, tt := range tester.testcases {
		testname := fmt.Sprintf("%dv,%d", tt.vertices, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := greedy(tester.logger, tt.graph)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
