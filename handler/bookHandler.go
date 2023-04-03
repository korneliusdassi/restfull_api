package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/korneliusdassi/restfull_api/model"
)

type bookHandler struct {
	bookService model.Service
}

func NewBookHandler(bookService model.Service) *bookHandler {
	return &bookHandler{bookService}
}

func responseBook(b model.Book) model.BookResponse {
	return model.BookResponse{
		Id:          b.Id,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Rating:      b.Rating,
	}
}

// get all data
func (handler *bookHandler) GetBooksHandler(c *gin.Context) {
	books, err := handler.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	var booksResponse []model.BookResponse
	for _, b := range books {
		bookResponse := responseBook(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": "Successfully",
		"data":   booksResponse,
	})
}

// get data by id
func (handler *bookHandler) GetBookHandler(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	b, err := handler.bookService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := responseBook(b)

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": "Successfully",
		"data":   bookResponse,
	})
}

// func method post
func (handler *bookHandler) PostBooksHandler(c *gin.Context) {
	var bookRequest model.BookRequest
	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {

		errorMessages := []string{} //untuk menampilkan banyak pesan error
		for _, error := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("field %s tidak boleh kosong, condition: %s", error.Field(), error.ActualTag())
			errorMessages = append(errorMessages, errMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	//panggil service
	book, err := handler.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": "Successfully",
		"data":   responseBook(book),
	})
}

func (handler *bookHandler) UpdateBooksHandler(c *gin.Context) {
	var bookRequest model.BookRequest
	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {

		errorMessages := []string{} //untuk menampilkan banyak pesan error
		for _, error := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("field %s tidak boleh kosong, condition: %s", error.Field(), error.ActualTag())
			errorMessages = append(errorMessages, errMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	//panggil service
	book, err := handler.bookService.Update(id, bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": "Successfully",
		"data":   responseBook(book),
	})
}

func (handler *bookHandler) DeleteBookHandler(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	_, err := handler.bookService.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	// bookResponse := responseBook(b)

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": "Successfully",
	})
}
