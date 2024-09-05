package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lahnasti/clientServerApp/server/handlers/jwt"
	"github.com/lahnasti/clientServerApp/server/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) LoginUserHandler(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid params", "error": err.Error()})
		return
	}

	// Проверяем наличие пользователя в базе
	userFromDB, err := s.Db.GetUserByLogin(user.Login)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials: user not found", "error": err.Error()})
		return
	}

	// Проверяем пароль
	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials: incorrect password", "error": err.Error()})
		return
	}

	uidStr := strconv.Itoa(userFromDB.UID)
	token, err := jwt.GenerateToken(uidStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate JWT token", "error": err.Error()})
		return
	}

	// Возвращаем токен в теле ответа
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
