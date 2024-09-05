package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url, filename, token string) error {
	// Создаем запрос
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/download?filename=%s", url, filename), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token) // Добавляем JWT-токен в заголовок

	// Отправляем запрос
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
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
	// Открываем файл для записи
	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer out.Close()

	// Копируем содержимое ответа в файл
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	fmt.Println("File downloaded successfully")
	return nil
}
