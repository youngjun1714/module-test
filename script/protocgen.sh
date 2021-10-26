#!/usr/bin/env bash

set -eo pipefail

proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)

for dir in $proto_dirs; do
    protoc \
        -I "proto" \
        --gocosmos_out=plugins=interfacetype+grpc,\
Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:. \
        $(find "${dir}" -maxdepth 1 -name '*.proto')

    protoc \
    -I "proto" \
    --grpc-gateway_out=logtostderr=true:. \
    $(find "${dir}" -maxdepth 1 -name '*.proto')

done 

cp -r github.com/youngjun1714/module-test/* ./
rm -rf github.com
