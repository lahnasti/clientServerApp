# clientServerApp
Тестовое задание: Создать простое клиент-серверное приложение для хранилища данных, в котором пользователи смогут загружать данные на сервер

## ТРЕБОВАНИЯ
- Go 1.22 или выше
- Postman (опционально, для тестирования API)
- Docker (опциаонально, для сборки приложения)

## ЗАПУСК ПРИЛОЖЕНИЯ БЕЗ ИСПОЛЬЗОВАНИЯ DOCKER
*Сборка сервера:*
1. Перейдите в директорию server/cmd: `cd clientServerApp/server/cmd`
2. Соберите серверную часть: `go build -o server main.go`
3. Запустите в первом терминале: `./main`
После этого сервер начнет прослушивать HTTP-запросы на порту 8080.

*Сборка клиента:*
1. Перейдите в директорию client: `cd clientServerApp/client/cmd`
2. Соберите клиентскую часть: `go build -o client main.go`
3. В зависимости от того, что вы хотите сделать (загрузить или скачать файл), выполните соответствующую команду во втором терминале:
- для загрузки файла`./client --upload /your/path/yourfile.txt`
- для скачивания файла `./client --download yourfile.txt`

## ИСПОЛЬЗОВАНИЕ API
*API предоставляет два основных маршрута для работы с файлами:* 
- загрузка (POST /upload)
- скачивание (GET /download)

*Загрузка файла (POST /upload)*
URL: http://localhost:8080/upload

Метод: POST

Описание: Загружает файл на сервер.

Параметры:

- file (формат данных form-data): Файл, который нужно загрузить.

*Пример с использованием curl:*
`curl -F "file=@path/to/your/file.txt" http://localhost:8080/upload`

*Пример с использованием Postman:*

Создайте новый запрос POST.
Введите URL http://localhost:8080/upload.
На вкладке Body выберите form-data.
В поле ключа добавьте file и выберите тип File.
Выберите файл и нажмите Send.

*Скачивание файла (GET /download)*
URL: http://localhost:8080/download

Метод: GET

Описание: Скачивает файл с сервера.

Параметры:

- filename (Query): Имя файла, который нужно скачать.

*Пример с использованием curl:*
`curl -O http://localhost:8080/download?filename=file.txt`

*Пример с использованием Postman:*

Создайте новый запрос GET.
Введите URL http://localhost:8080/download?filename=file.txt.
Нажмите Send и скачайте файл.

## ЗАПУСК ПРИЛОЖЕНИЯ С ПОМОЩЬЮ DOCKER
Введите в терминале `docker compose up --build`