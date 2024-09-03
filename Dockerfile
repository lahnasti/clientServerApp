# Используем официальный образ Go
FROM golang:1.22.4

# Устанавливаем рабочий каталог
WORKDIR /app

# Копируем все файлы проекта в контейнер
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект в рабочий каталог
COPY . .

# Сборка сервера
RUN go build -o srv ./server/cmd/main.go

# Сборка клиента
RUN go build -o cli ./client/cmd/main.go

# Определяем команду по умолчанию
CMD [ "./srv" ]
