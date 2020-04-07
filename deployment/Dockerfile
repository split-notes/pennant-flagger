FROM golang:1.13

WORKDIR /app

RUN go get github.com/githubnemo/CompileDaemon

## This container exposes port 8080 to the outside world
EXPOSE 7777

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -v -mod vendor -o .sprite ./" -command="./.sprite"
