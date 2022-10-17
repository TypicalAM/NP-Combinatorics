"""
Create a connected weighted graph for a Taxi Driver solution.
"""
import argparse
import json
import random
import time
import logging
from typing import Any, Callable

import networkx as nx
from networkx.readwrite import json_graph

MAX_SIZE = 200
DEFAULT_PATH = "example.json"


class CustomFormatter(logging.Formatter):
    """Custom logger formatting with colors"""

    green = "\033[32m"
    red = "\033[31m"
    blue = "\033[34m"
    gray = "\033[37m"
    reset = "\033[0m"
    levelname = "[%(levelname)s] "
    time = "%(asctime)s "
    msg = "%(message)s"

    FORMATS = {
        logging.INFO: green
        + levelname
        + reset
        + blue
        + time
        + reset
        + gray
        + msg
        + reset,
        logging.ERROR: red
        + levelname
        + reset
        + blue
        + time
        + reset
        + gray
        + msg
        + reset,
    }

    def format(self, record: Any) -> str:
        """Format the record and add coloring"""
        log_fmt = self.FORMATS.get(record.levelno)
        formatter = logging.Formatter(log_fmt)
        return formatter.format(record)


def get_logger() -> logging.Logger:
    """Get the colored logger"""
    logger = logging.getLogger("generate")
    logger.setLevel(logging.DEBUG)
    stream = logging.StreamHandler()
    stream.setLevel(logging.DEBUG)
    stream.setFormatter(CustomFormatter())
    logger.addHandler(stream)
    return logger


def timing(func: Callable) -> Callable:
    """Decorator to measure the time of a function run"""

    def wrapper(*arg, **kw) -> tuple[float, Any]:
        start = time.time()
        res = func(*arg, **kw)
        end = time.time() - start
        return (end, res)

    return wrapper


def parse_cmdline() -> argparse.Namespace:
    """Parse the commandline arguments provided to the app"""
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "--size",
        type=int,
        default=10,
        choices=range(1, MAX_SIZE),
        help=f"""Generate a graph of size n up to {MAX_SIZE} (default is 10)""",
        metavar="size",
    )
    parser.add_argument(
        "--path",
        default=DEFAULT_PATH,
        help=f"""State the path where the output should be saved (default is {DEFAULT_PATH})""",
        metavar="path",
    )
    parser.add_argument(
        "--density",
        type=int,
        default=100,
        choices=range(1, 101),
        help="""Change the density of the graph to n (default is 100)""",
        metavar="density",
    )
    parser.add_argument(
        "--min_weight",
        type=int,
        default=1,
        help="""Change the minimum weight of an edge of the graph to n (default is 1)""",
        metavar="min_weight",
    )
    parser.add_argument(
        "--max_weight",
        type=int,
        default=50,
        help="""Change the maximum weight of an edge of the graph to n (default is 50)""",
        metavar="min_weight",
    )
    return parser.parse_args()


@timing
def generate_connected_graph(args: argparse.Namespace) -> Any:
    """Return a  graph of size n"""
    graph = nx.erdos_renyi_graph(args.size, args.density / 100)
    for initial, target in graph.edges():
        graph.edges[initial, target]["weight"] = random.randint(args.min_weight, args.max_weight)
    return graph


def write_to_file(graph: Any, path: str) -> bool:
    """Dump the json data into the specified file"""
    try:
        file = open(path, "w", encoding="utf-8")
    except OSError:
        return False
    with file:
        graph_data = json_graph.adjacency_data(graph)
        json_object = json.dumps(graph_data, indent=2)
        file.write(json_object)
    return True


def main() -> None:
    """Driver code"""
    args = parse_cmdline()
    logger = get_logger()
    logger.info(
        "Generating a graph with size: %d, density %d, min %d, max %d",
        args.size, args.density, args.min_weight, args.max_weight
    )
    gentime, graph = generate_connected_graph(args)
    logger.info("The graph has been generated, it took %f seconds", gentime)
    if write_to_file(graph, args.path):
        logger.info("Graph dumped successfully to %s",args.path)
    else:
        logger.error("Encountered an error when writing to file %s", args.path)


if __name__ == "__main__":
    main()
