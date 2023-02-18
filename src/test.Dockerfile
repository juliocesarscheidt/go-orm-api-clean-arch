FROM golang:1.18-alpine as builder

WORKDIR /go/src/app

COPY go.mod go.sum /go/src/app/
RUN go mod download

COPY ./ ./

ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV CGO_ENABLED=0

ENTRYPOINT []
CMD ["/bin/sh", "-c", "go test tests/**/**/*_test.go -v"]
