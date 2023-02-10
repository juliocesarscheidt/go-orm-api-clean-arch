FROM golang:1.18-alpine as builder

WORKDIR /go/src/app

COPY go.mod go.sum /go/src/app/
RUN go mod download

COPY ./ ./

ENTRYPOINT []
CMD ["/bin/sh", "-c", "go test tests/**/**/*_test.go -v"]
