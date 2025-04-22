FROM golang:1.24-alpine AS builder

# Munkakönyvtár beállítása
WORKDIR /app

# Go mod és sum fájlok másolása
COPY go.mod go.sum ./

# Függőségek letöltése
RUN go mod download

# Forráskód másolása
COPY . .

# Alkalmazás fordítása
RUN CGO_ENABLED=0 GOOS=linux go build -o bookstore .

# Futtatási fázis
FROM alpine:latest

# Munkakönyvtár beállítása
WORKDIR /root/

# Konfigurációs fájl másolása
COPY --from=builder /app/config.yml .

# Fordított bináris másolása
COPY --from=builder /app/bookstore .

# Exponált port
EXPOSE 8080

# Indítási parancs
CMD ["./bookstore"]