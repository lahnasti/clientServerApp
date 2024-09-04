package client

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func UploadFile(filename string) error {
	fmt.Println("Trying to open file:", filename)
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", filename, err)
	}
	defer file.Close()

	fmt.Println("File opened successfully")

	// Создаем поток данных
	bodyReader, bodyWriter := io.Pipe()
	writer := multipart.NewWriter(bodyWriter)

	go func() {
		defer bodyWriter.Close()
		part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
		if err != nil {
			fmt.Println("Error creating form file:", err)
			return
		}

		// Потоковое копирование данных из файла в форму multipart.Writer
		_, err = io.Copy(part, file)
		if err != nil {
			fmt.Println("Error copying file to part:", err)
			return
		}

		// Закрываем writer, чтобы завершить формирование данных
		writer.Close()
	}()

	// Отправляем POST-запрос на сервер
	req, err := http.NewRequest("POST", "http://server:8080/upload", bodyReader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Проверяем успешность ответа
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to upload file: %s", resp.Status)
	}

	fmt.Println("File uploaded successfully")
	return nil
}
func DownloadFile(filename string) error {
	// Запрос на скачивание
	resp, err := http.Get("http://server:8080/download?filename=" + filename)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: %s", resp.Status)
	}

	// Читаем содержимое ответа
	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, resp.Body)
	if err != nil {
		return err
	}

	// Выводим содержимое на консоль
	fmt.Println("File contents:")
	fmt.Println(buffer.String())
	// Открытие файла для записи
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	fmt.Println("File downloaded successfully")
	return nil
}
