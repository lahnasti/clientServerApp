package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func DownloadFileHandler (ctx *gin.Context) {
	filename := ctx.Query("filename")
	if filename == "" {
		ctx.String(http.StatusBadRequest, "Filename is required")
        return
	}
	filePath := filepath.Join(uploadPath, filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		ctx.String(http.StatusNotFound, "File not found")
        return
	}
	ctx.File(filePath)
}