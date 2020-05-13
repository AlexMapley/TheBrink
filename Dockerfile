FROM golang

COPY ./ /go/src/the_brink

ENTRYPOINT /go/src/entrypoint.sh