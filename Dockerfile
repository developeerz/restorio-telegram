FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o restorio-telegram-service ./cmd/restorio-telegram


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/restorio-telegram-service /app/

CMD ["./restorio-telegram-service"]