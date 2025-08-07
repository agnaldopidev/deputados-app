# Etapa de build
FROM golang:1.22.2 AS builder

ARG SERVICE=rest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/${SERVICE}/main.go

# Etapa final
FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 8080

ENTRYPOINT ["/app/app"]
