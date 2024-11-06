FROM golang:1.23

# Создаем директорию приложения и копируем файлы
RUN mkdir /app
ADD . /app
WORKDIR /app

# Сборка приложения
RUN go build -o main cmd/main.go

# Устанавливаем точку входа
ENTRYPOINT ["/app/main"]