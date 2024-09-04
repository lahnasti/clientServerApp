# Используем официальный образ Go
FROM golang:1.22.4

# Устанавливаем рабочий каталог
WORKDIR /app

# Копируем файлы go модулей и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект в рабочий каталог
COPY . .

# Сборка сервера и клиента
RUN go build -o srv ./server/cmd/main.go
RUN go build -o cli ./client/cmd/main.go

# Определяем команду по умолчанию для сервера
CMD [ "./start.sh" ]