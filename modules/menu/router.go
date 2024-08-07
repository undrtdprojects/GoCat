package book

import (
	"quiz-3-sanbercode-greg/helpers/common"
	"quiz-3-sanbercode-greg/middlewares"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api/categories/:id")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("/books", CreateBookRouter)
		api.GET("/books", GetAllBookRouter)
		api.GET("/books/:id", GetBookByIdRouter)
		api.PUT("/books/:id", UpdateBookRouter)
		api.DELETE("/books/:id", DeleteBookRouter)
	}
}

func CreateBookRouter(ctx *gin.Context) {
	var (
		bookRepo = NewRepository()
		bookSrv  = NewService(bookRepo)
	)

	_, err := bookSrv.CreateBookService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully added book data")
}

func GetAllBookRouter(ctx *gin.Context) {
	var (
		bookRepo = NewRepository()
		bookSrv  = NewService(bookRepo)
	)

	books, err := bookSrv.GetAllBookService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get all book data", books)
}

func GetBookByIdRouter(ctx *gin.Context) {
	var (
		bookRepo = NewRepository()
		bookSrv  = NewService(bookRepo)
	)

	books, err := bookSrv.GetBookByIdService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get book data", books)
}

func DeleteBookRouter(ctx *gin.Context) {
	var (
		bookRepo = NewRepository()
		bookSrv  = NewService(bookRepo)
	)

	err := bookSrv.DeleteBookService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully delete book data")
}

func UpdateBookRouter(ctx *gin.Context) {
	var (
		bookRepo = NewRepository()
		bookSrv  = NewService(bookRepo)
	)

	err := bookSrv.UpdateBookService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully update book data")
}
