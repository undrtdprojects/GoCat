package book

import (
	"errors"
	"fmt"
	"quiz-3-sanbercode-greg/helpers/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateBookService(ctx *gin.Context) (err error)
	GetAllBookService(ctx *gin.Context) (result []Book, err error)
	GetBookByIdService(ctx *gin.Context) (result Book, err error)
	DeleteBookService(ctx *gin.Context) (err error)
	UpdateBookService(ctx *gin.Context) (err error)
}

type bookService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &bookService{repository}
}

func (service *bookService) CreateBookService(ctx *gin.Context) (err error) {
	var newBook Book

	err = ctx.ShouldBind(&newBook)
	if err != nil {
		return err
	}

	var categories []Book
	categories, err = service.repository.GetAllBookRepository()
	if err != nil {
		return err
	}

	book, err := service.repository.GetBookByTitleRepository(newBook.Title)
	if err != nil {
		return err
	}

	if len(categories) != 0 && book.Title != "" {
		err = errors.New("book already exists")
		return err
	}

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField()

	newBook.CreatedAt = defaultField.CreatedAt
	newBook.CreatedBy = defaultField.CreatedBy
	newBook.ModifiedAt = defaultField.ModifiedAt
	newBook.ModifiedBy = defaultField.ModifiedBy

	fmt.Println("create categories :", newBook)
	err = service.repository.CreateBookRepository(newBook)
	if err != nil {
		return errors.New("failed to add new book")
	}

	return
}

func (service *bookService) GetAllBookService(ctx *gin.Context) (categories []Book, err error) {
	return service.repository.GetAllBookRepository()
}

func (service *bookService) GetBookByIdService(ctx *gin.Context) (book Book, err error) {
	var (
		idInt int
		id    = ctx.Param("id")
	)

	idInt, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id book from param")
		return
	}

	return service.repository.GetBookByIdRepository(idInt)
}

func (service *bookService) GetAllBooksByBookService(ctx *gin.Context) (books []book.Book, err error) {
	var (
		idInt int
		id    = ctx.Param("id")
		name  = ctx.Param("name")
	)

	idInt, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id book from param")
		return
	}

	return service.repository.GetAllBooksByBookRepository(idInt, name)
}

func (service *bookService) DeleteBookService(ctx *gin.Context) (err error) {
	var (
		book Book
		id   = ctx.Param("id")
	)

	book.Id, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id book from param")
		return
	}

	return service.repository.DeleteBookRepository(book)
}

func (service *bookService) UpdateBookService(ctx *gin.Context) (err error) {
	var (
		book Book
		id   = ctx.Param("id")
	)

	err = ctx.ShouldBind(&book)
	if err != nil {
		return
	}

	book.Id, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id book from param")
		return
	}
	return service.repository.UpdateBookRepository(book)
}
