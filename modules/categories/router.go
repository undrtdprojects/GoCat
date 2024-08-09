package categories

import (
	"GoCat/databases/connection"
	"GoCat/helpers/common"
	"GoCat/middlewares"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	api.Use(middlewares.RoleCheck())
	{
		api.POST("/categories", CreateCategoriesRouter)
		api.GET("/categories", GetAllCategoriesRouter)
		api.GET("/categories/:id", GetCategoriesByIdRouter)
		api.PUT("/categories/:id", UpdateCategoriesRouter)
		api.DELETE("/categories/:id", DeleteCategoriesRouter)
	}
}

func CreateCategoriesRouter(ctx *gin.Context) {
	var (
		categoriesRepo = NewRepository(connection.DBConnections)
		categoriesSrv  = NewService(categoriesRepo)
	)

	err := categoriesSrv.CreateCategoriesService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully added categories data")
}

func GetAllCategoriesRouter(ctx *gin.Context) {
	var (
		categoriesRepo = NewRepository(connection.DBConnections)
		categoriesSrv  = NewService(categoriesRepo)
	)

	categories, err := categoriesSrv.GetAllCategoriesService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithListData(ctx, "successfully get all categories data", int64(len(categories)), categories)
}

func GetCategoriesByIdRouter(ctx *gin.Context) {
	var (
		categoriesRepo = NewRepository(connection.DBConnections)
		categoriesSrv  = NewService(categoriesRepo)
	)

	categories, err := categoriesSrv.GetCategoriesByIdService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get categories data", categories)
}

func DeleteCategoriesRouter(ctx *gin.Context) {
	var (
		categoriesRepo = NewRepository(connection.DBConnections)
		categoriesSrv  = NewService(categoriesRepo)
	)

	err := categoriesSrv.DeleteCategoriesService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully delete categories data")
}

func UpdateCategoriesRouter(ctx *gin.Context) {
	var (
		categoriesRepo = NewRepository(connection.DBConnections)
		categoriesSrv  = NewService(categoriesRepo)
	)

	err := categoriesSrv.UpdateCategoriesService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully update categories data")
}
