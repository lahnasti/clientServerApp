package repository

type Repository interface {
	GetFile(string) error
	UploadFile(string)error
}