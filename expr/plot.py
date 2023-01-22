#!/usr/bin/env python3


import argparse
import csv

import matplotlib.patches as mpatches
import numpy as np
from matplotlib import pyplot as plt
from matplotlib.path import Path

COLUMN_HOST = "host"
COLUMN_OVERALL_LATENCY = "overallDuration"


def main(log_file, stats_file, log_figure_file, stats_figure_file):
    plot_log(log_file, log_figure_file)
    plot_stats(stats_file, stats_figure_file)


def plot_log(file, figure_file):
    rows = []
    with open(file, "r") as f:
        reader = csv.DictReader(f)

        rows = list(reader)

    d = []
    for r in rows:
        x = int(r[COLUMN_OVERALL_LATENCY]) / 1000
        d.append(x)

    fig, ax = plt.subplots(
        nrows=2,
        figsize=(3, 4),
        dpi=160,
        sharex="col",
        gridspec_kw={"height_ratios": (1, 2)},
    )

    # plot two same graph
    ax[0].boxplot(d)
    ax[1].boxplot(d, sym="")
    # print(d)
    ax[1].set_ylabel("latency (ms)")

    # kuttukeru
    fig.subplots_adjust(hspace=0.0)

    # ticker (lower side)
    ax[1].set_ylim(5, 8)
    ax[1].set_yticks(np.arange(5, 8, 1))

    # ticker (upper side)
    ax[0].set_ylim(1310, 1370)
    ax[0].set_yticks(np.arange(1310, 1370 + 1, 20))

    # hide notches
    ax[1].spines["top"].set_visible(False)
    ax[0].spines["bottom"].set_visible(False)
    ax[0].tick_params(axis="x", which="both", bottom=False, labelbottom=False)

    # nyoro nyoro
    d1 = 0.02  # X軸のはみだし量
    d2 = 0.03  # ニョロ波の高さ
    wn = 21  # ニョロ波の数（奇数値を指定）

    pp = (0, d2, 0, -d2)
    px = np.linspace(-d1, 1 + d1, wn)
    py = np.array([1 + pp[i % 4] for i in range(0, wn)])
    p = Path(list(zip(px, py)), [Path.MOVETO] + [Path.CURVE3] * (wn - 1))

    line1 = mpatches.PathPatch(
        p,
        lw=4,
        edgecolor="black",
        facecolor="None",
        clip_on=False,
        transform=ax[1].transAxes,
        zorder=10,
    )

    line2 = mpatches.PathPatch(
        p,
        lw=3,
        edgecolor="white",
        facecolor="None",
        clip_on=False,
        transform=ax[1].transAxes,
        zorder=10,
        capstyle="round",
    )

    _ = ax[1].add_patch(line1)
    _ = ax[1].add_patch(line2)

    plt.tight_layout()
    plt.savefig(figure_file)
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

    parser.add_argument(
        "log_figure_file",
        metavar="LOG_FIGURE_FILE",
        type=str,
        help="file name to save the graph of log",
    )
    parser.add_argument(
        "stats_figure_file",
        metavar="STATS_FIGURE_FILE",
        type=str,
        help="file name to save the graph of stats",
    )

    args = parser.parse_args()
    main(
        log_file=args.log_file,
        stats_file=args.stats_file,
        log_figure_file=args.log_figure_file,
        stats_figure_file=args.stats_figure_file,
    )
