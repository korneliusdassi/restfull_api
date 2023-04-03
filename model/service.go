package model

type Service interface {
	FindAll() ([]Book, error)
	FindById(id int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Update(id int, bookRequest BookRequest) (Book, error)
	Delete(id int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (servis *service) FindAll() ([]Book, error) {
	books, err := servis.repository.FindAll()
	return books, err
	// return s.repository.FindAll()
}

func (servis *service) FindById(id int) (Book, error) {
	return servis.repository.FindById(id)
}

func (servis *service) Create(bookRequest BookRequest) (Book, error) {
	book := Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Price:       bookRequest.Price,
		Rating:      bookRequest.Rating,
	}

	newBook, err := servis.repository.Create(book)
	return newBook, err
}

func (servis *service) Update(id int, bookRequest BookRequest) (Book, error) {
	book, _ := servis.repository.FindById(id)

	book.Title = bookRequest.Title
	book.Description = bookRequest.Description
	book.Price = bookRequest.Price
	book.Rating = bookRequest.Rating

	newBook, err := servis.repository.Update(book)
	return newBook, err
}

func (servis *service) Delete(id int) (Book, error) {
	book, _ := servis.repository.FindById(id)
	newBook, err := servis.repository.Delete(book)
	return newBook, err
}
