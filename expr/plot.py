#!/usr/bin/env python3

import argparse
import csv

from matplotlib import pyplot as plt
import numpy as np


COLUMN_HOST = "host"
COLUMN_OVERALL_LATENCY = "overallDuration"


def main(log_file, stats_file):
    plot_log(log_file)
    plot_stats(stats_file)


def plot_log(file):
    rows = []
    with open(file, "r") as f:
        reader = csv.DictReader(f)

        rows = list(reader)

    d = []
    for r in rows:
        x = int(r[COLUMN_OVERALL_LATENCY])
        d.append(x)

    fig = plt.figure()

    ax = fig.add_subplot(2, 1, 1)
    ax.figure.set_figwidth(10.0)
    ax.figure.set_figheight(10.0)
    ax.set_title("latencies")
    ax.set_ylabel("latency (ms)")

    ax.boxplot(d)

    plt.tight_layout()
    plt.show()


def plot_stats(file):
    pass


def group_by(column, rows):
    d = {}

    for i, r in enumerate(rows):
        x = r[column]
        if d.get(x) is None:
            d[x] = []

        r["originalIndex"] = i
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
