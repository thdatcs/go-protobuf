#!/bin/bash
GOPATH=${PWD}

cd src
for pkg in $@
do
    echo "golangci-lint $pkg"
    sources=$(GOPATH=${GOPATH} go list $pkg/...)
    GOPATH=${GOPATH} ${GOPATH}/bin/golangci-lint run ${sources}
done
cd ..