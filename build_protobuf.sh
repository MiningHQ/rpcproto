#!/bin/bash
function generate_protobuf() {
  protoc -I. --go_out=plugins=grpc:. ./$1/*.proto
}

echo 'Generating rpcproto/*'
for path in *
do
  if [ -d $path ] ; then
    generate_protobuf $path
  fi
done
