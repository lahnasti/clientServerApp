#!/bin/bash

# Запускаем контейнеры
docker-compose up -d

# Регистрация пользователя
docker-compose exec client ./cli -register -username nastya -password nastya -url http://server:8080

# Логин и получение токена
TOKEN=$(docker-compose exec client ./cli -login -username testuser -password testpassword -url http://server:8080 | grep 'Token:' | awk '{print $2}')

# Проверяем, что токен был получен
if [ -z "$TOKEN" ]; then
  echo "Failed to obtain token"
  exit 1
fi

# Загрузка файла
docker-compose exec client ./cli -upload /app/file.txt -token $TOKEN -url http://server:8080
