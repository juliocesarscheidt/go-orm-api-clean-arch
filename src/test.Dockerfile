FROM golang:1.21-alpine as builder
LABEL maintainer="Julio Cesar <julio@blackdevs.com.br>"

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV CGO_ENABLED=0

ENTRYPOINT []
CMD ["/bin/sh", "-c", "go test tests/application/**/*_test.go -v"]
