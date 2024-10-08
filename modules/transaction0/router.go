package transaction0

import (
	"GoCat/databases/connection"
	"GoCat/helpers/common"
	"GoCat/middlewares"
	"GoCat/modules/menu"
	"GoCat/modules/payment"
	"GoCat/modules/transaction1"
	"GoCat/modules/user"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	api.Use(middlewares.RoleCheck())
	{
		api.POST("/transaction0", CreateTransaction0Router)
		api.GET("/transaction0s", GetAllTransaction0Router)
		api.GET("/transaction0/:id", GetTransaction0ByIdRouter)
		api.PUT("/transaction0/:id", UpdateTransaction0Router)
		api.DELETE("/transaction0/:id", DeleteTransaction0Router)
	}
}

func CreateTransaction0Router(ctx *gin.Context) {
	var (
		transaction0Repo = NewRepository(connection.DBConnections)
		userRepo         = user.NewRepository(connection.DBConnections)
		paymentRepo      = payment.NewRepository(connection.DBConnections)
		menuRepo         = menu.NewRepository(connection.DBConnections)
		trans1Repo       = transaction1.NewRepository(connection.DBConnections)
		transaction0Srv  = NewService(transaction0Repo, userRepo, paymentRepo, menuRepo, trans1Repo)
	)

	err := transaction0Srv.CreateTransaction0Service(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully added transaction0 data")
}

func GetAllTransaction0Router(ctx *gin.Context) {
	var (
		transaction0Repo = NewRepository(connection.DBConnections)
		userRepo         = user.NewRepository(connection.DBConnections)
		paymentRepo      = payment.NewRepository(connection.DBConnections)
		menuRepo         = menu.NewRepository(connection.DBConnections)
		trans1Repo       = transaction1.NewRepository(connection.DBConnections)
		transaction0Srv  = NewService(transaction0Repo, userRepo, paymentRepo, menuRepo, trans1Repo)
	)

	transaction0s, err := transaction0Srv.GetAllTransaction0Service(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get all transaction0 data", transaction0s)
}

func GetTransaction0ByIdRouter(ctx *gin.Context) {
	var (
		transaction0Repo = NewRepository(connection.DBConnections)
		userRepo         = user.NewRepository(connection.DBConnections)
		paymentRepo      = payment.NewRepository(connection.DBConnections)
		menuRepo         = menu.NewRepository(connection.DBConnections)
		trans1Repo       = transaction1.NewRepository(connection.DBConnections)
		transaction0Srv  = NewService(transaction0Repo, userRepo, paymentRepo, menuRepo, trans1Repo)
	)

	transaction0, err := transaction0Srv.GetTransaction0ByIdService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get transaction0 data", transaction0)
}

func DeleteTransaction0Router(ctx *gin.Context) {
	var (
		transaction0Repo = NewRepository(connection.DBConnections)
		userRepo         = user.NewRepository(connection.DBConnections)
		paymentRepo      = payment.NewRepository(connection.DBConnections)
		menuRepo         = menu.NewRepository(connection.DBConnections)
		trans1Repo       = transaction1.NewRepository(connection.DBConnections)
		transaction0Srv  = NewService(transaction0Repo, userRepo, paymentRepo, menuRepo, trans1Repo)
	)

	err := transaction0Srv.DeleteTransaction0Service(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully delete transaction0 data")
}

func UpdateTransaction0Router(ctx *gin.Context) {
	var (
		transaction0Repo = NewRepository(connection.DBConnections)
		userRepo         = user.NewRepository(connection.DBConnections)
		paymentRepo      = payment.NewRepository(connection.DBConnections)
		menuRepo         = menu.NewRepository(connection.DBConnections)
		trans1Repo       = transaction1.NewRepository(connection.DBConnections)
		transaction0Srv  = NewService(transaction0Repo, userRepo, paymentRepo, menuRepo, trans1Repo)
	)

	err := transaction0Srv.UpdateTransaction0Service(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully update transaction0 data")
}
