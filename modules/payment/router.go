package payment

import (
	"GoCat/databases/connection"
	"GoCat/helpers/common"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api")
	// api.Use(middlewares.JwtMiddleware())
	// api.Use(middlewares.Logging())
	{
		api.POST("/payment", CreatePaymentRouter)
		api.GET("/payments", GetAllPaymentRouter)
		api.GET("/payment/:id", GetPaymentByIdRouter)
		api.PUT("/payment/:id", UpdatePaymentRouter)
		api.DELETE("/payment/:id", DeletePaymentRouter)
	}
}

func CreatePaymentRouter(ctx *gin.Context) {
	var (
		paymentRepo = NewRepository(connection.DBConnections)
		paymentSrv  = NewService(paymentRepo)
	)

	err := paymentSrv.CreatePaymentService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully added payment data")
}

func GetAllPaymentRouter(ctx *gin.Context) {
	var (
		paymentRepo = NewRepository(connection.DBConnections)
		paymentSrv  = NewService(paymentRepo)
	)

	payments, err := paymentSrv.GetAllPaymentService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get all payment data", payments)
}

func GetPaymentByIdRouter(ctx *gin.Context) {
	var (
		paymentRepo = NewRepository(connection.DBConnections)
		paymentSrv  = NewService(paymentRepo)
	)

	payment, err := paymentSrv.GetPaymentByIdService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get payment data", payment)
}

func DeletePaymentRouter(ctx *gin.Context) {
	var (
		paymentRepo = NewRepository(connection.DBConnections)
		paymentSrv  = NewService(paymentRepo)
	)

	err := paymentSrv.DeletePaymentService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully delete payment data")
}

func UpdatePaymentRouter(ctx *gin.Context) {
	var (
		paymentRepo = NewRepository(connection.DBConnections)
		paymentSrv  = NewService(paymentRepo)
	)

	err := paymentSrv.UpdatePaymentService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully update payment data")
}
