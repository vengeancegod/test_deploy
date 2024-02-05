# Stage 1: Сборка приложения
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Финальный образ
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]
