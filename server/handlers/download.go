package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func DownloadFileHandler(c *gin.Context) {
	// Получаем имя файла из параметра запроса
	filename := c.Query("filename")
	filepath := "./uploads/" + filename

	// Открываем файл
	file, err := os.Open(filepath)
	if err != nil {
		c.String(http.StatusNotFound, fmt.Sprintf("file not found: %s", err.Error()))
		return
	}
	defer file.Close()

	// Устанавливаем заголовки для скачивания файла
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.Writer.Header().Add("Content-Length", fmt.Sprintf("%d", getFileSize(filepath)))

	// Передаем файл на скачивание
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("download file err: %s", err.Error()))
		return
	}
}


// Вспомогательная функция для получения размера файла
func getFileSize(filepath string) int64 {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}