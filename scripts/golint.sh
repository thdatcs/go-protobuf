#!/bin/bash
GOPATH=${PWD}

cd src
for pkg in $@
do
    echo "golint $pkg"
    sources=$(GOPATH=${GOPATH} go list $pkg/...)
    GOPATH=${GOPATH} ${GOPATH}/bin/golint ${sources}
done
cd ..