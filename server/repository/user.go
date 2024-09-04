package repository

import "github.com/lahnasti/clientServerApp/server/models"

type UserRepo interface {
	RegisterUser(models.User)(int, error)
	LoginUser(models.User)(int, error)
}