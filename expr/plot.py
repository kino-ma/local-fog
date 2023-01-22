#!/usr/bin/env python3

import argparse
import csv

from matplotlib import pyplot as plt


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

    total_ax = fig.add_subplot(n // 2 + 1, 2, n + 1)
    total_ax.set_title("total")
    total_ax.set_xlabel("latency (ms)")
    total_ax.set_ylabel("times of request")
    total_ax.set_yscale("log")

    for i, (host, latcs) in enumerate(sorted(latencies.items())):
        print(f"{host}: {len(latcs)} rows ([0] = {latcs[0]}")
        i = i + 1
        ax = fig.add_subplot(n // 2 + 1, 2, i)
        ax.set_yscale("log")
        ax.set_ylim(5, 2000)
        ax.figure.set_figwidth(10.0)
        ax.figure.set_figheight(10.0)
        ax.set_title(host)
        ax.set_ylabel("latency (ms)")
        ax.scatter(range(1, len(latcs) + 1), latcs)
        ax.set_aspect(1.0 / ax.get_data_ratio())

        total_ax.scatter(
            range(1, len(latcs) + 1),
            latcs,
            label=host,
        )

    latcs = []
    for ls in latencies.values():
        latcs.extend(ls)

    plt.tight_layout()
    plt.show()


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
