# syntax=docker/dockerfile:1

# Build

FROM golang:1.18 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY *.go .

RUN go build -o /webhdv03

# Deploy

FROM debian:latest

WORKDIR /

COPY --from=build /webhdv03 /usr/local/bin/webhdv03

EXPOSE 8080

ENTRYPOINT [ "/usr/local/bin/webhdv03" ]