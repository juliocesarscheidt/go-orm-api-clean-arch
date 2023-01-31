FROM golang:1.18-alpine as builder
LABEL maintainer="Julio Cesar <julio@blackdevs.com.br>"

WORKDIR /go/src/app

COPY go.mod go.sum /go/src/app/
RUN go mod download

COPY ./ ./

RUN GOOS=linux GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 \
    go build -ldflags="-s -w" -o /go/src/app/main .

FROM gcr.io/distroless/static:nonroot

WORKDIR /
COPY --from=builder /go/src/app/main .
USER nonroot:nonroot

EXPOSE 8000

ENTRYPOINT [ "/main" ]
