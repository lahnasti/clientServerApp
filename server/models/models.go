package models

type User struct {
	UID int `json:"uid"` 
	Login string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}