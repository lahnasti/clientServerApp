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
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return err
	}
	
	// Копируем содержимое файла в форму
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	// Закрываем writer, чтобы завершить формирование данных
	err = writer.Close()
	if err != nil {
		return err
	}

	// Отправляем POST-запрос на сервер
	resp, err := http.Post("http://localhost:8080/upload", writer.FormDataContentType(), body)
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
	resp, err := http.Get("http://localhost:8080/download?filename=" + filename)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: %s", resp.Status)
	}

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("File downloaded successfully")
	return nil
}
