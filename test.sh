#!/bin/bash

# Запускаем контейнеры
docker-compose up -d

# Регистрация пользователя
docker-compose exec client ./cli -register -login user -password test -url http://server:8080

# Логин и получение токена
TOKEN=$(docker-compose exec client ./cli -auth -login user -password test -url http://server:8080 | grep 'Token:' | awk '{print $2}')

# Проверяем, что токен был получен
if [ -z "$TOKEN" ]; then
  echo "Failed to obtain token"
  exit 1
fi

# Копирование файла на контейнер
docker cp /your/path/file.txt clientserverapp-client-1:/app/file.txt

# Загрузка файла
docker-compose exec client ./cli -upload /app/file.txt -token $TOKEN -url http://server:8080

# Проверяем, что файл был загружен успешно

# Скачивание файла
docker-compose exec client ./cli -download file.txt -token $TOKEN -url http://server:8080

# Проверяем, что файл был скачан успешно

# Остановка и удаление контейнеров
docker-compose down
