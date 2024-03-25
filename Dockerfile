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

ENV REDIS_ADDR "redis:6379"
ENV REDIS_PASSWORD ""
ENV REDIS_DB 0
ENV REDIS_TIMER 15

ENV FLOOD_TIME_DIF_MIL 6420
ENV FLOOD_MAX_REQUEST 3
ENV FLOOD_COUNT_REQUEST 25
ENV FLOOD_TIME_SLEEP_MIL 2000

CMD ./main





