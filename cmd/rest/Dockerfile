FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o rest-api ./cmd/rest

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/rest-api .

EXPOSE 8080
CMD ["./rest-api"]