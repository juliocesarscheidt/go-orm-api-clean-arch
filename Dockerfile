FROM golang:1.18-alpine as builder

WORKDIR /go/src/app

COPY go.mod go.sum /go/src/app/
RUN go mod download

COPY ./ ./

RUN GOOS=linux GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 \
    go build -ldflags="-s -w" -o /go/src/app/main .

FROM gcr.io/distroless/static:nonroot

LABEL maintainer="Julio Cesar <julio@blackdevs.com.br>"
LABEL org.opencontainers.image.source https://github.com/juliocesarscheidt/go-orm-api-clean-arch
LABEL org.opencontainers.image.description "Simple Go API implementing clean architecture, using Mux, Go ORM and Prometheus"
LABEL org.opencontainers.image.licenses=MIT

WORKDIR /
COPY --from=builder /go/src/app/main .
USER nonroot:nonroot

EXPOSE 8000

ENTRYPOINT [ "/main" ]
