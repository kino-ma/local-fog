#!/usr/bin/env bash

echo 'starting experiment...'

stdbuf -oL docker stats --format '{{.Name}},{{.CPUPerc}},{{.MemPerc}},{{.NetIO}}' > /tmp/stats.log &
stats_pid=$!

docker compose up -d node1 node2 node3 cloud
docker compose run --rm client &>/dev/null

echo 'client container finished. shutting down...'

docker compose down
kill $stats_pid
wait $stats_pid

sed -u 's/\x1b\[[0-9;]*[mGKHFJ]//g' /tmp/stats.log > log/stats.csv
rm /tmp/stats.log

echo 'plot figures...'
./expr/plot.py log/log.csv log/stats.csv log/figs/log.png log/figs/stats.png

echo 'done.'