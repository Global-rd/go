# Stage 1: Build the Go binary
FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o webservice .

FROM debian:bookworm-slim
RUN apt-get update && apt-get -y install procps net-tools
WORKDIR /app/
COPY --from=builder /app/webservice .
COPY movie-ws.yaml /app/movie-ws.yaml
CMD ["./webservice"]
