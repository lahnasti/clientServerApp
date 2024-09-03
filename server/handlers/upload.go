package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const uploadPath = "./upload"

func UploadFileHandler(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, "Failed to get file: %s", err.Error())
		return
	}
	err = os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to create upload directory: %s", err.Error())
		return
	}
	filePath := filepath.Join(uploadPath, filepath.Base(file.Filename))
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to save file: %s", err.Error())
        return
	}
	ctx.String(http.StatusOK, "File uploaded successfully: %s", file.Filename)
}