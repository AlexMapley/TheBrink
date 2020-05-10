ls
cd /go/src/the_brink

go clean
go install ./...
pwd
ls
echo

go build  
./the_brink


# export GOBIN=$GOPATH/bin
# go build ./... main
# GOPATH=~/
# export GOBIN=$GOPATH/bin