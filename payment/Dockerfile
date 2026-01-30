# syntax=docker/dockerfile:1

FROM golang:latest AS builder
WORKDIR /src

# Necessario para os 'replace ../../microservices-proto/...' no go.mod
COPY microservices-proto ./microservices-proto

# Copia apenas o microsservico alvo
COPY microservices/payment ./microservices/payment

WORKDIR /src/microservices/payment
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -trimpath -ldflags="-s -w" -o /out/payment ./cmd

FROM gcr.io/distroless/base-debian12:nonroot
COPY --from=builder /out/payment /payment

EXPOSE 3001
ENTRYPOINT ["/payment"]