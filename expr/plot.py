#!/usr/bin/env python3


import argparse
import csv
import re

import matplotlib.patches as mpatches
import numpy as np
from matplotlib import pyplot as plt
from matplotlib.path import Path

COLUMN_HOST = "host"
COLUMN_OVERALL_LATENCY = "overallDuration"

COLUMN_CONTAINER_NAME = "NAME"
COLUMN_NET_IO = "NET I/O"

UNIT_KILO = 1_000
UNIT_MEGA = 1_000 * UNIT_KILO
UNIT_GIGA = 1_000 * UNIT_MEGA


def main(log_file, stats_file, log_figure_file, stats_figure_file):
    # plot_log(log_file, log_figure_file)
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


def plot_stats(file, figure_file):
    rows = []
    with open(file, "r") as f:
        reader = csv.DictReader(f)

        rows = list(reader)

    groups = group_by(COLUMN_CONTAINER_NAME, rows)

    if groups.get("--"):
        del groups["--"]

    d_out = []
    d_in = []

    for host, rows in sorted(groups.items()):
        print(f"host = {host}")

        i, o = get_netio(rows)
        ii = np.array(i, dtype=float)
        oo = np.array(o, dtype=float)
        print("i, o", ii.shape, oo.shape)
        d_in.append(ii)
        d_out.append(oo)

    print("shape of out, in")
    print(np.shape(d_out))
    print(np.shape(d_in))

    shortest = min(map(lambda a: a.shape[0], d_out))
    print(f"{shortest=}")

    for i, o in zip(d_in, d_out):
        i.resize(shortest, refcheck=False)
        o.resize(shortest, refcheck=False)

    print("shape of out, in")
    print(np.shape(d_out))
    print(np.shape(d_in))

    d_out = np.vstack(d_out)
    d_in = np.vstack(d_in)

    fig = plt.figure()
    ax_out = fig.add_subplot(2, 1, 1)
    ax_out.stackplot(np.arange(0, d_out.shape[1]), d_out)
    ax_in = fig.add_subplot(2, 1, 2)
    ax_in.stackplot(np.arange(0, d_in.shape[1]), d_in)

    ax_out.set_ylabel("bytes sent (GB)")
    ax_in.set_ylabel("bytes received (GB)")
    ax_out.set_xlabel("# of requests")
    ax_in.set_xlabel("# of requests")

    ax_out.legend(
        ["cloud", "node A", "node B", "node C", "client"],
        loc="upper left",
    )
    ax_in.legend(
        ["cloud", "node A", "node B", "node C", "client"],
        loc="upper left",
    )

    plt.tight_layout()
    plt.savefig(figure_file)
    plt.show()


def get_netio(rows):
    i = []
    o = []
    p = re.compile("([0-9.]+[kMG]?B) / ([0-9.]+[kMG]?B)")

    for row in rows:
        io = row[COLUMN_NET_IO]
        if not io:
            continue

        matches = p.search(io)
        ii, oo = matches.groups()

        in_bytes, out_bytes = size_to_byte_int(ii), size_to_byte_int(oo)

        i.append(in_bytes), o.append(out_bytes)

    return i, o


def size_to_byte_int(size_str):
    last2 = size_str[-2:]
    last1 = size_str[-1:]
    if last2 == "GB":
        x = float(size_str[:-2])
        return x * UNIT_GIGA
    elif last2 == "MB":
        x = float(size_str[:-2])
        return x * UNIT_MEGA
    elif last2 == "kB":
        x = float(size_str[:-2])
        return x * UNIT_KILO
    elif last1 == "B":
        x = float(size_str[:-1])
        return x
    else:
        raise RuntimeError(f"invalid size string '{size_str}'")


def container_to_name(container_name):
    p = re.compile("local-fog_([a-z0-9]*)_.*")
    matches = p.search(container_name)
    return matches and matches.groups(1)


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
