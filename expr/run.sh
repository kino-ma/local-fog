#!/usr/bin/env bash

echo 'starting experiment...'

program=$0

usage() {
    echo "$program identifier"
}

identifier=$1

if [[ -z $identifier ]]
then
    usage
    exit 1
fi

log="log/$identifier-log.csv"
tmp_stats="/tmp/$identifier-stats.log"
stats="log/$identifier-stats.csv"

stdbuf -oL docker stats --format 'table {{.Name}},{{.CPUPerc}},{{.MemPerc}},{{.NetIO}}' > "$tmp_stats" &
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
./expr/plot.py log/log.csv log/stats.csv log/figs/log.png log/figs/stats.png

echo 'done.'