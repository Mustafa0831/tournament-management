FROM golang:1.22.6-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем только необходимые файлы
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Переходим в директорию сервиса
WORKDIR /app/api-gateway

# Собираем бинарный файл
RUN go build -o main .

# Команда для запуска\
CMD ["./main"]
