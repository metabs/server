FROM golang:1.13.4 AS builder
ENV GO111MODULE="on"
ENV GOFLAGS=" -mod=vendor"
WORKDIR /go/src/github.com/unprogettosenzanomecheforseinizieremo/server
COPY . .
RUN make

FROM ubuntu:bionic-20190424
RUN apt-get update && apt-get install -y apt-utils ca-certificates
COPY --from=builder /go/src/github.com/unprogettosenzanomecheforseinizieremo/server/service-account.json/ /service-account.json
COPY --from=builder /go/src/github.com/unprogettosenzanomecheforseinizieremo/server/dist/ /
