FROM golang:1.21-alpine AS builder
WORKDIR /build
COPY . .
RUN go build ./cmd/main.go

FROM ubuntu:latest

ENV DEBIAN_FRONTEND=noninteractive

USER root

WORKDIR /rest
COPY --from=builder /build/main .

COPY . .


CMD ./main





