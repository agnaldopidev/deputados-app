# Stage 1: Build
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o grpc-server ./cmd/grpc

# Stage 2: Run
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/grpc-server .

EXPOSE 50051

CMD ["./grpc-server"]
