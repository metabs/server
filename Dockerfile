FROM golang:1.14.1 AS builder
WORKDIR /go/src/github.com/metabs/server
COPY . .
RUN make

FROM ubuntu:bionic-20190424
RUN apt-get update && apt-get install -y apt-utils ca-certificates
COPY --from=builder /go/src/github.com/metabs/server/service-account.json/ /service-account.json
COPY --from=builder /go/src/github.com/metabs/server/dist/ /
