FROM golang:latest

COPY ./ /go/src/the_brink

RUN chmod +x /go/src/the_brink/entrypoint.sh

ENTRYPOINT /go/src/the_brink/entrypoint.sh