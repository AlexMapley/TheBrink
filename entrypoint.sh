#!/bin/sh

ls
cd /go/src/the_brink

go clean
go install ./...
ls

go build
./the_brink
