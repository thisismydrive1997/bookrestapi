package repoimpl

import (
	"example/restfulapi/models"
	"example/restfulapi/repository"
	"fmt"

	"gorm.io/gorm"
)

type BookRepoImpl struct {
	Db *gorm.DB
}

func NewBookRepo(db *gorm.DB) repository.BookRepo {
	return &BookRepoImpl{
		Db: db,
	}
}

func (b *BookRepoImpl) FindBookAll() ([]models.Book, error) {
	fmt.Println("FindBookAll")
	books := []models.Book{}

	err := b.Db.Table(models.TableBooks).Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *BookRepoImpl) FindBookById(id string) (models.Book, error) {
	fmt.Println("FindBookById:", id)
	book := models.Book{}

	err := b.Db.Table(models.TableBooks).Where("id=?", id).Find(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (b *BookRepoImpl) CreateBook(book models.Book) error {
	fmt.Println("Create book")

	err := b.Db.Table(models.TableBooks).Create(&book).Error
	if err != nil {
		return err
	}

	return nil
}

func (b *BookRepoImpl) UpdateBook(book models.Book) error {
	fmt.Println("Update book")

	err := b.Db.Table(models.TableBooks).Where("id=?", book.Id).Updates(book).Error
	if err != nil {
		return err
	}

	return nil
}

func (b *BookRepoImpl) DeleteBookById(id string) error {
	fmt.Println("Delete book id:", id)

	err := b.Db.Table(models.TableBooks).Delete(&models.Book{Id: id}).Error
	if err != nil {
		return err
	}

	return nil
}
