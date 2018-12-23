#!/bin/bash
function clean_protobuf() {
  rm ./$1/*.pb.go
}

for path in *
do
  if [ -d $path ] ; then
    clean_protobuf $path
  fi
done
