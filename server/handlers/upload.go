package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadFileHandler(c *gin.Context) {
	// Создаем папку uploads, если она не существует
	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		err := os.Mkdir("./uploads", os.ModePerm)
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to create directory: %v", err))
			return
		}
	}

	// Получаем файл из запроса
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error retrieving file: %s", err.Error()))
		return
	}
	defer file.Close()

	fmt.Printf("Received file: %s\n", header.Filename) // Добавлено для отладки

	// Создаем файл на диске в папке uploads
	out, err := os.Create("./uploads/" + header.Filename)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error creating file: %s", err.Error()))
		return
	}
	defer out.Close()

	// Копируем данные из входящего файла в файл на диске блоками
	_, err = io.Copy(out, file)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error copying file: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", header.Filename))
}
