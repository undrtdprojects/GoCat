package categories

import (
	"quiz-3-sanbercode-greg/databases/connection"
	"quiz-3-sanbercode-greg/helpers/common"
	"quiz-3-sanbercode-greg/middlewares"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.POST("/categories", CreateCategoryRouter)
		api.GET("/categories", GetAllCategoryRouter)
		api.GET("/categories/:id", GetCategoryByIdRouter)
		api.GET("/categories/:id/books", GetAllBooksByCategoryRouter)
		api.PUT("/categories/:id", UpdateCategoryRouter)
		api.DELETE("/categories/:id", DeleteCategoryRouter)
	}
}

func CreateCategoryRouter(ctx *gin.Context) {
	var (
		categoryRepo = NewRepository(connection.DBConnections)
		categorySrv  = NewService(categoryRepo)
	)

	err := categorySrv.CreateCategoryService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully added category data")
}

func GetAllCategoryRouter(ctx *gin.Context) {
	var (
		categoryRepo = NewRepository(connection.DBConnections)
		categorySrv  = NewService(categoryRepo)
	)

	categories, err := categorySrv.GetAllCategoryService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithListData(ctx, "successfully get all category data", int64(len(categories)), categories)
}

func GetCategoryByIdRouter(ctx *gin.Context) {
	var (
		categoryRepo = NewRepository(connection.DBConnections)
		categorySrv  = NewService(categoryRepo)
	)

	category, err := categorySrv.GetCategoryByIdService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get category data", category)
}

func GetAllBooksByCategoryRouter(ctx *gin.Context) {
	var (
		categoryRepo = NewRepository(connection.DBConnections)
		categorySrv  = NewService(categoryRepo)
	)

	category, err := categorySrv.GetCategoryByIdService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get category data", category)
}

func DeleteCategoryRouter(ctx *gin.Context) {
	var (
		categoryRepo = NewRepository(connection.DBConnections)
		categorySrv  = NewService(categoryRepo)
	)

	err := categorySrv.DeleteCategoryService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully delete category data")
}

func UpdateCategoryRouter(ctx *gin.Context) {
	var (
		categoryRepo = NewRepository(connection.DBConnections)
		categorySrv  = NewService(categoryRepo)
	)

	err := categorySrv.UpdateCategoryService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully update category data")
}
