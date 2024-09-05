package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lahnasti/clientServerApp/server/handlers"
	"github.com/lahnasti/clientServerApp/server/handlers/jwt"
)

func SetupRoutes(s *handlers.Server) *gin.Engine {
	r := gin.Default()

	// Маршруты для аутентификации
	authGroup := r.Group("/")
	{
		authGroup.POST("/register", s.RegisterUserHandler)
		authGroup.POST("/login", s.LoginUserHandler)
	}

	// Маршруты для пользователей с применением JWT Middleware
	userGroup := r.Group("/")
	userGroup.Use(jwt.JWTAuthMiddleware()) // Применяем middleware
	{
		userGroup.POST("/upload", handlers.UploadFileHandler)
		userGroup.GET("/download/", handlers.DownloadFileHandler)
	}

	return r
}