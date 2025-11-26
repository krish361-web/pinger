# Этап 1: Сборка (Builder)
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем файлы
COPY go.mod ./
COPY main.go ./

# Компилируем в статический бинарник (CGO_ENABLED=0 отвязывает от системных библиотек)
# -ldflags="-s -w" удаляет отладочную инфу, уменьшая размер файла
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server main.go

# Этап 2: Финальный образ (Runner)
FROM alpine:latest

# Устанавливаем ping (iputils)
RUN apk add --no-cache iputils

WORKDIR /root/

# Копируем только скомпилированный файл из первого этапа
COPY --from=builder /app/server .

EXPOSE 80

CMD ["./server"]