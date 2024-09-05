package client

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func UploadFile(url, filename, token string) error {
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
	req, err := http.NewRequest("POST", url+"/upload", bodyReader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Проверяем успешность ответа
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to upload file: %s, response body: %s", resp.Status, string(body))
	}

	fmt.Println("File uploaded successfully")
	return nil
}
