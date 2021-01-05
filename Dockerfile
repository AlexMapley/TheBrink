FROM golang:latest

COPY ./ /go/src/the_brink

RUN chmod +x /go/src/the_brink/entrypoint.sh

RUN go get -d -v /go/src/the_brink/...
RUN go install -v /go/src/the_brink/...

ENTRYPOINT /go/src/the_brink/entrypoint.sh