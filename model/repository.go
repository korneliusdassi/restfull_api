package model

import "gorm.io/gorm"

type Repository interface {
	Create(book Book) (Book, error)
	FindAll() ([]Book, error)
	FindById(id int) (Book, error)
	Update(book Book) (Book, error)
	Delete(book Book) (Book, error)
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

func (r *RepositoryImpl) FindAll() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *RepositoryImpl) FindById(id int) (Book, error) {
	var book Book
	err := r.db.Find(&book, id).Error
	return book, err
}

func (r *RepositoryImpl) Create(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *RepositoryImpl) Update(book Book) (Book, error) {
	err := r.db.Save(&book).Error
	return book, err
}

func (r *RepositoryImpl) Delete(book Book) (Book, error) {
	err := r.db.Delete(&book).Error
	return book, err
}
