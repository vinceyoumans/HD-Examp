# syntax=docker/dockerfile:1

# Build

FROM golang:1.18 AS build

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY *.go .

RUN go build -o /hellov02

# Deploy

FROM scratch

WORKDIR /

COPY --from=build /hellov02 /hellov02

CMD [ "/hellov02" ]