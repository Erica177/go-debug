FROM golang:1.18 AS build

COPY . /work

RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /work

RUN go mod tidy && STATIC=0 GOOS=linux GOARCH=amd64 LDFLAGS='-extldflags -static -s -w' go build -o main

FROM ubuntu:xenial

COPY --from=build /work/main /main
COPY --from=build /work/debug/cmd/viewcore/viewcore /viewcore

RUN apt-get update && apt-get install -y gdb

ENTRYPOINT ["./main"]