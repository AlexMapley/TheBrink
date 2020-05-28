# Creating A Game Container

### What you'll need
1. A working code editor is helpful
    - Vscode is my favorite: https://code.visualstudio.com/download
    - Sublime is good as well: https://www.sublimetext.com/
    - Atom is also good: https://atom.io/
2. docker
    - If you need to install docker, check out https://docs.docker.com/get-docker/

## Creating your base program
For now, let's just keep it simple.

The first file we will add to our game is just `main.go`, as follows:

```golang
package main


func main() {
    // The beginning
    var yourGame := "default video game title"
	fmt.Printf("Welcome to your game: %s\n", yourGame)
}
```

## Setting up your console
We're going to be running your videogame out of a `docker` container capabale of running `golang`.

The best docker container for this is argubaly `golang:latest`, which is regularly updated
and kept up to the latest version of go.

In the same directory as your `main.go`, add another file titled `Dockerfile` with the following:

```docker
FROM golang:latest

COPY ./ /go/src/my_rpg

RUN chmod +x /go/src/my_rpg/entrypoint.sh

ENTRYPOINT /go/src/my_rpg/entrypoint.sh
```

In the same direcotry again, you can also add `entrypoint.sh`:

```sh
#!/bin/sh

cd /go/src/my_rpg

go clean
go install ./...

go build
./my_rpg
```