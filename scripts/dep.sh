#!/bin/bash
GOPATH=${PWD}

for pkg in $@
do
    echo "dep $pkg"
    cd $GOPATH/src/$pkg 
    if [ ! -f "$GOPATH/src/$pkg/Gopkg.toml" ]; then
        GOPATH=${GOPATH} $GOPATH/bin/dep init 
    fi
    GOPATH=${GOPATH} $GOPATH/bin/dep ensure -v
    cd $GOPATH
done