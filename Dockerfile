FROM golang:1.13.4 AS builder
WORKDIR /go/src/github.com/damianopetrungaro/app
COPY . .
RUN make

FROM ubuntu:bionic-20190424
RUN apt-get update && apt-get install -y apt-utils ca-certificates
COPY --from=builder /go/src/github.com/damianopetrungaro/app/out/ /
