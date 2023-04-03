package main

import (
	"github.com/gin-gonic/gin"

	"github.com/korneliusdassi/restfull_api/handler"
	"github.com/korneliusdassi/restfull_api/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//melakukan koneksi ke db
	dsn := "root:@tcp(127.0.0.1:3306)/restfull_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// otomatis membuat tabel
	db.AutoMigrate(&model.Book{})

	bookRepository := model.NewRepository(db)
	bookService := model.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	r := gin.Default()
	r.GET("/books/:id", bookHandler.GetBookHandler) //get data by id
	r.GET("/books", bookHandler.GetBooksHandler)    //get all data
	r.POST("/books", bookHandler.PostBooksHandler)
	r.PUT("/books/:id", bookHandler.UpdateBooksHandler)
	r.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	r.Run("localhost:3000") //menjalankan server
}
