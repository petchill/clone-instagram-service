FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /clone-instagram-service

FROM alpine:latest

WORKDIR /app

COPY --from=builder /clone-instagram-service .

EXPOSE 5000

CMD ["./clone-instagram-service"]