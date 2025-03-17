FROM golang:1.24-alpine AS builder
WORKDIR /app

# Go mod dosyalarını kopyala
COPY go.mod go.sum ./

# GOPROXY ayarını yapılandır (önce proxy, sonra direkt)
ENV GOPROXY=https://proxy.golang.org,direct

# Modülleri indir
RUN go mod download

# Uygulama dosyalarını kopyala ve build et
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o food-api ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/food-api .
COPY --from=builder /app/configs ./configs
CMD ["./food-api"]
