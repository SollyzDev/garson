#!/bin/sh
echo "go path is" $GOPATH
echo $(pwd)
echo $(ls code)
p=$GOPATH/src/github.com/ishehata/garson
mkdir -p $p
cp -R code/* $p/
cd $p
echo "verify that code files are copied"
echo "pwd is $(pwd)"
echo $(ls)
echo "installing dependencies"
go get
echo "prepairing to start go test"
go vet .
go test -v .
