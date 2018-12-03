#!/usr/bin/env bash

set -ex

go clean -cache
git clean -xdf

cd src
./make.bash
cd ..

for i in {1..50}
do
    ./bin/go run ../git-bisect-run/crash2.go
done
