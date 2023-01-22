#!/usr/bin/env bash

program=$0

usage() {
    printf "usage:\t$program identifier\n"
}

identifier=$1

if [[ -z $identifier ]]
then
    usage
    exit 1
fi

echo 'starting experiment...'

log="log/$identifier-log.csv"
log_figure="log/figs/$identifier-log.png"
tmp_stats="/tmp/$identifier-stats.log"
stats="log/$identifier-stats.csv"
stats_figure="log/figs/$identifier-stats.png"


echo 'NAME,CPU %,MEM %,NET I/O' > "$tmp_stats"
stdbuf -oL docker stats --format '{{.Name}},{{.CPUPerc}},{{.MemPerc}},{{.NetIO}}' >> "$tmp_stats" &
stats_pid=$!

docker compose up -d node1 node2 node3 cloud
docker compose run --rm client "$log" &>/dev/null

echo 'client container finished. shutting down...'

docker compose down
kill $stats_pid
wait $stats_pid

sed -u 's/\x1b\[[0-9;]*[mGKHFJ]//g' $tmp_stats > $stats
rm $tmp_stats

echo 'plot figures...'
./expr/plot.py "$log" "$stats" "$log_figure" "$stats_figure"

echo 'done.'