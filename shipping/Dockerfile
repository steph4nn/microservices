# syntax=docker/dockerfile:1

FROM golang:latest AS builder
WORKDIR /src

# Necessario para o 'replace ../../microservices-proto/...' no go.mod
COPY microservices-proto ./microservices-proto

# Copia apenas o microsservico alvo
COPY microservices/shipping ./microservices/shipping

WORKDIR /src/microservices/shipping
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/shipping ./cmd

FROM gcr.io/distroless/base-debian12:nonroot
COPY --from=builder /out/shipping /shipping

EXPOSE 3002
ENTRYPOINT ["/shipping"]
