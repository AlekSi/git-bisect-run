#!/usr/bin/env bash

set -ex

go clean -cache
git clean -xdf

cd src
./make.bash
cd ..

for i in {1..1000}
do
    ./bin/go run ../git-bisect-run/crash.go
done
