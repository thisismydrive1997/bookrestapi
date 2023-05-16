package repository

import "example/restfulapi/models"

type BookRepo interface {
	FindBookAll() ([]models.Book, error)
	FindBookById(string) (models.Book, error)
	CreateBook(models.Book) error
	UpdateBook(models.Book) error
	DeleteBookById(string) error
}
