FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum .

RUN go mod download

COPY . . 

RUN go build -o backend ./cmd/app

FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/backend .

COPY --from=builder /app/internal/migrations ./migrations

CMD ["./backend"]

