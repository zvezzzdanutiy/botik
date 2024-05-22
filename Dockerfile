# Устанавливаем базовый образ Golang
FROM --platform=linux/amd64 golang:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /newbot

# Копируем все файлы из локального проекта "newbot" в рабочую директорию контейнера
COPY newbot .

# Устанавливаем зависимости
RUN go mod download

# Собираем приложение
RUN go build cmd/main.go

# Запускаем приложение при запуске контейнера
CMD ["./main"]
