# Stage Build
FROM golang:1.24.1-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .

# Stage Run
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080

ENV LOG_LEVEL=debug
ENV SERVER_PORT=8080

CMD ["sh", "-c", "./main"]