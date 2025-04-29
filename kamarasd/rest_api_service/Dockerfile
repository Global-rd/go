FROM golang:1.23.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o rest_api_service .

FROM gcr.io/distroless/static:nonroot

WORKDIR /app

COPY --from=builder /app/rest_api_service /app/rest_api_service 

EXPOSE 8080