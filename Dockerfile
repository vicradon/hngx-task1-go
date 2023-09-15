FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .



FROM alpine:latest AS releaser

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
