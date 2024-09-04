package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lahnasti/clientServerApp/server/handlers"
)

func UserRoutes() *gin.Engine {
	r := gin.Default()
	userGroup := r.Group("/")
	{
		userGroup.POST("/upload", handlers.UploadFileHandler)
		userGroup.GET("/download/", handlers.DownloadFileHandler)
	}
	return r
}
