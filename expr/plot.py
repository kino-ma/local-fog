#!/usr/bin/env python3

import argparse
import csv
from pprint import pprint

from matplotlib import pyplot as plt


def main(log_file, stats_file):
    plot_log(log_file)
    plot_stats(stats_file)


def plot_log(file):
    rows = []
    with open(file, "r") as f:
        reader = csv.DictReader(f)

        rows = list(reader)

    groups = group_by("host", rows)

    pprint(groups)


def plot_stats(file):
    pass


def group_by(column, rows):
    d = {}

    for r in rows:
        x = r[column]
        if d.get(x) is None:
            d[x] = []

        d[x].append(r)

    return d


if __name__ == "__main__":
    parser = argparse.ArgumentParser(prog="plot.py", description="Plot graph")
    parser.add_argument(
        "log_file",
        metavar="LOG_FILE",
        type=str,
        help="file name of the log file",
    )
    parser.add_argument(
        "stats_file",
        metavar="STATS_FILE",
        type=str,
        help="file name of the stats file",
    )

    args = parser.parse_args()
    main(
        log_file=args.log_file,
        stats_file=args.stats_file,
    )
