package transaction1

import (
	"GoCat/databases/connection"
	"GoCat/helpers/common"
	"GoCat/modules/menu"
	"GoCat/modules/transaction0"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api")
	// api.Use(middlewares.JwtMiddleware())
	// api.Use(middlewares.Logging())
	{
		api.POST("/transaction1", CreateTransaction1Router)
		api.GET("/transaction1", GetAllTransaction1Router)
		api.GET("/transaction1/:id", GetTransaction1ByIdRouter)
		api.PUT("/transaction1/:id", UpdateTransaction1Router)
		api.DELETE("/transaction1/:id", DeleteTransaction1Router)
	}
}

func CreateTransaction1Router(ctx *gin.Context) {
	var (
		transaction1Repo = NewRepository(connection.DBConnections)
		transaction0Repo = transaction0.NewRepository(connection.DBConnections)
		menuRepo         = menu.NewRepository(connection.DBConnections)
		transaction1Srv  = NewService(transaction1Repo, transaction0Repo, menuRepo)
	)

	err := transaction1Srv.CreateTransaction1Service(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully added transaction1 data")
}

func GetAllTransaction1Router(ctx *gin.Context) {
	var (
		transaction1Repo = NewRepository(connection.DBConnections)
		transaction0Repo = transaction0.NewRepository(connection.DBConnections)
		menuRepo         = menu.NewRepository(connection.DBConnections)
		transaction1Srv  = NewService(transaction1Repo, transaction0Repo, menuRepo)
	)

	transaction1s, err := transaction1Srv.GetAllTransaction1Service(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get all transaction1 data", transaction1s)
}

func GetTransaction1ByIdRouter(ctx *gin.Context) {
	var (
		transaction1Repo = NewRepository(connection.DBConnections)
		transaction0Repo = transaction0.NewRepository(connection.DBConnections)
		menuRepo         = menu.NewRepository(connection.DBConnections)
		transaction1Srv  = NewService(transaction1Repo, transaction0Repo, menuRepo)
	)

	transaction1, err := transaction1Srv.GetTransaction1ByIdService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get transaction1 data", transaction1)
}

func DeleteTransaction1Router(ctx *gin.Context) {
	var (
		transaction1Repo = NewRepository(connection.DBConnections)
		transaction0Repo = transaction0.NewRepository(connection.DBConnections)
		menuRepo         = menu.NewRepository(connection.DBConnections)
		transaction1Srv  = NewService(transaction1Repo, transaction0Repo, menuRepo)
	)

	err := transaction1Srv.DeleteTransaction1Service(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully delete transaction1 data")
}

func UpdateTransaction1Router(ctx *gin.Context) {
	var (
		transaction1Repo = NewRepository(connection.DBConnections)
		transaction0Repo = transaction0.NewRepository(connection.DBConnections)
		menuRepo         = menu.NewRepository(connection.DBConnections)
		transaction1Srv  = NewService(transaction1Repo, transaction0Repo, menuRepo)
	)

	err := transaction1Srv.UpdateTransaction1Service(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully update transaction1 data")
}
