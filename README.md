# Combinatorial Optimization Project (NP-Complete problem solver)

A program by Adam Piaseczny (L12) which allows the user to
- Generate a weighted graph using `python3`
- Solve the TSP problem using `golang`

## What is the TSP problem?

The travelling salesman problem asks the following question: "Given a list of cities and the distances between each pair of cities, what is the shortest possible route that visits each city exactly once and returns to the origin city?". This can be formally expressed as finding the shortest hamiltonian circuit in a complete weighted graph.

## What are the implemented strategies

The implemented solvers are the following
- Bruteforce
	- The bruteforce approach generates all possible permutations of the set of vertices ($n!$) and then computes the total distance of every one of them, keeping the smallest one in a variable
- Greedy
	- The greedy solver always takes the nearest unvisited vertex and "jumps to it" until all the vertices are exhausted, this allows for a fast approximation for a rather expensive problem.
- Backtracking
	- Check every path, but when the path is bigger than an already discovered minimal path, discard it

## Usage of the program

The generator is written in `python3` while the `solver` is written in `golang`. First, clone the repository:

```bash
git clone https://github.com/TypicalAM/NP-Combinatorics/ && cd NP-Combinatorics
```

### Generator (`python3`)

To use the generator (on a linux machine)

```bash
cd generator
# python3 -m pip install --user virtualenv
virtualenv --no-vcs-ignore venv
source venv/bin/activate
pip3 install -r requirements.txt
python3 generate.py --help
```

Now you can run the generator, for this projects complete graphs were used, so the density should stay at 100. An example generation of a graph using the erdos renyi model would look something like this:

```bash
python3 generate.py \
	--size 15 \
	--density 100 \
	--min_weight 20 \
	--max_weight 50 \
	--path ../solver/data/15_vertices.json
```

The generator will dump a `json` file looking a little bit like example.json.

### Solver (`golang`)

To use the solver on a linux machine (needs `go` to be installed)

```bash
cd solver
go build -o solver main.go
./solver --help
```

Examples of generation and solving could be

```bash
# in solver dir
python3 ../generator/generate.py --path ../solver/data/10_vertices.json
./solver --load data/10_vertices.json
```

```bash
# in solver dir
python3 ../generator/generate.py --path ../solver/data/10_vertices.json
./solver --load data/10_vertices.json
```

```bash
# in solver dir
python3 ../generator/generate.py \
	--size 30 \
	--density 100 \
	--min_weight 2 \
	--max_weight 15 \
	--path ../solver/data/30_vertices.json
./solver \
	--load data/30_vertices.json \
	--solvers greedy \
	--show=false
```

```bash
# in solver dir
python3 ../generator/generate.py \
	--size 10 \
	--min_weight 1 \
	--max_weight 200 \
	--path ../solver/data/10_vertices.json
./solver \
	--load data/10_vertices.json \
	--solvers bruteforce,backtracking \
	--show=true
```

## Testing (solver)

To test the solver run the following commands in the solver directory:

```bash
go test ./src/solutions
```

or to see a more verbose output

```bash
go test -v ./src/solutions
```
