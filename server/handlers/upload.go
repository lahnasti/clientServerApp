package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadFileHandler(ctx *gin.Context) {
	 // Создаем папку uploads, если она не существует
	 if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
        err := os.Mkdir("./uploads", os.ModePerm)
        if err != nil {
            log.Fatalf("Failed to create directory: %v", err)
        }
    }
	// Получаем файл из запроса
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	defer file.Close()

	// Создаем файл на диске
	out, err := os.Create("./uploads/" + header.Filename)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("create file err: %s", err.Error()))
		return
	}
	defer out.Close()

	// Копируем данные из входящего файла в файл на диске блоками
	_, err = io.Copy(out, file)
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("copy file err: %s", err.Error()))
		return
	}

	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", header.Filename))
}
