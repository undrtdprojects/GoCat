package menu

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
		api.POST("/menu", CreateMenuRouter)
		api.GET("/menus", GetAllMenuRouter)
		api.GET("/menu/:id", GetMenuByIdRouter)
		api.PUT("/menu/:id", UpdateMenuRouter)
		api.DELETE("/menu/:id", DeleteMenuRouter)
	}
}

func CreateMenuRouter(ctx *gin.Context) {
	var (
		menuRepo = NewRepository(connection.DBConnections)
		menuSrv  = NewService(menuRepo)
	)

	err := menuSrv.CreateMenuService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully added menu data")
}

func GetAllMenuRouter(ctx *gin.Context) {
	var (
		menuRepo = NewRepository(connection.DBConnections)
		menuSrv  = NewService(menuRepo)
	)

	menus, err := menuSrv.GetAllMenuService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get all menu data", menus)
}

func GetMenuByIdRouter(ctx *gin.Context) {
	var (
		menuRepo = NewRepository(connection.DBConnections)
		menuSrv  = NewService(menuRepo)
	)

	menu, err := menuSrv.GetMenuByIdService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get menu data", menu)
}

func DeleteMenuRouter(ctx *gin.Context) {
	var (
		menuRepo = NewRepository(connection.DBConnections)
		menuSrv  = NewService(menuRepo)
	)

	err := menuSrv.DeleteMenuService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully delete menu data")
}

func UpdateMenuRouter(ctx *gin.Context) {
	var (
		menuRepo = NewRepository(connection.DBConnections)
		menuSrv  = NewService(menuRepo)
	)

	err := menuSrv.UpdateMenuService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully update menu data")
}
