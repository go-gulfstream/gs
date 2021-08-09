#!/bin/bash
path="$(pwd)"
declare -a mocks=(
    "internal/stream/command_mutation"
    "internal/stream/event_mutation"
    "internal/projection/projection"
)
for i in "${mocks[@]}"
do
   parts=($(echo $i | tr '/' "\n"))
   index=$((${#parts[@]}-2))
   pkg="${parts[index]}"
   docker run --rm -v $path:$path github.com/go-gulfstream/mockgen -package="mock$pkg" -destination="$path/mocks/${i}.go" -source="$path/${i}.go"
done;