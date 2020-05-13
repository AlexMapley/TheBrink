FROM golang

COPY ./ /go/src/the_brink

RUN chmod +x /go/src/the_brink

ENTRYPOINT /go/src/the_brink/entrypoint.sh