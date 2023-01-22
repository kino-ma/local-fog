#!/usr/bin/env python3

import argparse
import csv

from matplotlib import pyplot as plt
import numpy as np
import matplotlib.patches as mpatches
from matplotlib.path import Path


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

    # kuttukeru
    fig.subplots_adjust(hspace=0.0)

    ax[1].set_ylim(5000, 8000)
    ax[1].set_yticks(np.arange(5000, 8000, 1000))

    ax[0].set_ylim(1310000, 1370000)
    ax[0].set_yticks(np.arange(1310000, 1370000 + 1, 20000))

    ax[1].spines["top"].set_visible(False)

    ax[0].spines["bottom"].set_visible(False)
    ax[0].tick_params(axis="x", which="both", bottom=False, labelbottom=False)

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

    a = ax[1].add_patch(line1)
    a = ax[1].add_patch(line2)

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
