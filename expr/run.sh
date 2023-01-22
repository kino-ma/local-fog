#!/usr/bin/env bash

docker stats --format '{{.Name}},{{.CPUPerc}},{{.MemPerc}},{{.NetIO}}' | sed 's/\x1b\[[0-9;]*[mGKHFJ]//g' > log/stats.csv &
stats_pid=$!

docker compose up -d node1 node2 node3 cloud
docker compose run --rm client
docker compose down
kill $stats_pid