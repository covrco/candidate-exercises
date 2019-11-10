#!/bin/sh
# gripmock runs on a linux system with some expectations, protoc, protoc-gen-go
# et cetera, its recommended to invoke this script inside the docker
# container that satisfies those constraints

set -e

/usr/bin/gripmock src/*.proto &
pid=$(ps -o pid,comm | grep gripmock | awk '{print $1}')

finish()
{
  kill -s SIGINT ${pid}
  echo "Stopped gripmock pid for grpc proto"
  exit
}

trap finish SIGINT

rpcs=$(ls mocks | grep -v README.md)

for rpc in ${rpcs}; do
  for stub in $(ls mocks/${rpc}); do
    curl -s -X POST -d @"mocks/${rpc}/${stub}" localhost:4771/add
    echo "Registered stub=${stub} for rpc=${rpc}."
  done
done

while :; do
  # docker logs --since 20s ${container_id}
  sleep 5
done

