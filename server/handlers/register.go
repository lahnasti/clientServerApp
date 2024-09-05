package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lahnasti/clientServerApp/server/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	RegisterUser(models.User) (int, error)
	GetUserByLogin(string) (models.User, error)
}

type Server struct {
	Db UserRepo
}

func NewServer(db UserRepo) *Server {
	return &Server{
		Db: db,
	}
}

func (s *Server) RegisterUserHandler(ctx *gin.Context) {
	// Регистрация пользователя
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid params", "error": err.Error()})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash password", "error": err.Error()})
		return
	}
	user.Password = string(hash)
	uid, err := s.Db.RegisterUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register user", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "uid": uid})
}
