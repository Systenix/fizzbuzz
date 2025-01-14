# Use official Golang image as builder
FROM golang:1.22.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fizzbuzz cmd/main.go

# Use a minimal base image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/fizzbuzz .

CMD ["./fizzbuzz"]