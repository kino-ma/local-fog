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

    groups = group_by(COLUMN_HOST, rows)

    latencies = {}
    for k, v in groups.items():
        ls = []
        for r in v:
            ls.append(int(r[COLUMN_OVERALL_LATENCY]) / 1000)

        latencies[k] = ls

    fig = plt.figure()

    n = len(latencies.keys())

    ax_node = fig.add_subplot(2, 1, 1)
    ax_node.figure.set_figwidth(10.0)
    ax_node.figure.set_figheight(10.0)
    ax_node.set_title("fog nodes")
    ax_node.set_ylabel("latency (ms)")
    ax_node.set_ylim(5.0, 11.0)

    ax_cloud = fig.add_subplot(2, 1, 2)
    ax_cloud.figure.set_figwidth(10.0)
    ax_cloud.figure.set_figheight(10.0)
    ax_cloud.set_title("cloud")
    ax_cloud.set_ylabel("latency (ms)")

    d_nodes = list(
        map(
            lambda x: x[1],
            filter(lambda item: item[0] != "cloud", latencies.items()),
        )
    )
    d_cloud = list(
        map(
            lambda x: x[1],
            filter(lambda item: item[0] == "cloud", latencies.items()),
        )
    )

    print(len(d_nodes))

    ax_node.boxplot(d_nodes, sym="")
    ax_cloud.boxplot(d_cloud)

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
