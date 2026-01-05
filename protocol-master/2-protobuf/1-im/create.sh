#!/bin/sh
SRC_DIR=./
DST_DIR=./

#C++
#mkdir -p $DST_DIR/cpp
#/usr/bin/protoc -I=$SRC_DIR --cpp_out=$DST_DIR $SRC_DIR/*.proto
protoc -I=$SRC_DIR --cpp_out=$DST_DIR $SRC_DIR/*.proto
#protoc -I=./ --cpp_out=./*.proto

#protoc -I=$SRC_DIR --java_out=$DST_DIR $SRC_DIR/*.proto