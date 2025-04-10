FROM golang:1.23.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/app /app/app

USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["/app/app"]