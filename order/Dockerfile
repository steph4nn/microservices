# syntax=docker/dockerfile:1

FROM golang:latest AS builder
WORKDIR /src

# Necessario para os 'replace ../../microservices-proto/...' nos go.mod
COPY microservices-proto ./microservices-proto

# Copia apenas o microsservico alvo
COPY microservices/order ./microservices/order

WORKDIR /src/microservices/order
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/order ./cmd

FROM gcr.io/distroless/base-debian12:nonroot
COPY --from=builder /out/order /order

EXPOSE 3000
ENTRYPOINT ["/order"]
