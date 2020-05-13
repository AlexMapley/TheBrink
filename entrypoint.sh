#!/bin/sh

ls
cd /go/src/the_brink

go clean
go install ./...
pwd
ls
echo

go build  
./the_brink
