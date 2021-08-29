FROM golang:latest
# Build
WORKDIR /work
COPY ./ /work

RUN go build -o the_brink
# Deploy
FROM alpine
RUN apk add --no-cache  libc6-compat
WORKDIR /game
COPY --from=0 /work/the_brink /game/
RUN chmod +x /game/the_brink
ENTRYPOINT [ "/game/the_brink" ]