#!/bin/bash

go get github.com/tools/godep
$GOPATH/bin/godep restore
go build -o bin/application application.go

