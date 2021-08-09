#!/bin/bash
ROOT_PATH=$(pwd)
for proto in $(find . -name *.proto);
do
    protopath=$ROOT_PATH$(dirname "${proto/\./}")
    echo $protopath
    docker run \
    --rm -v $ROOT_PATH:$ROOT_PATH -w $ROOT_PATH github.com/go-gulfstream/protoc:latest  \
     --proto_path=$protopath --go_out=Mgoogle/protobuf/timestamp.proto=github.com/google/protobuf/types,plugins=grpc:. -I . $proto
done;