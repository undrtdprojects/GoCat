package user

import (
	"GoCat/databases/connection"
	"GoCat/helpers/common"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api/users")
	{
		api.POST("/login", Login)
		api.POST("/signup", SignUp)
		api.PUT("/change-password", ChangePassword)
		api.GET("/users", GetListUser)
	}
}

func Login(ctx *gin.Context) {
	var (
		userRepo = NewRepository(connection.DBConnections)
		userSrv  = NewService(userRepo)
	)

	token, err := userSrv.LoginService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully login", token)
}

func SignUp(ctx *gin.Context) {
	var (
		userRepo = NewRepository(connection.DBConnections)
		userSrv  = NewService(userRepo)
	)

	err := userSrv.SignUpService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "awesome, successfully create user")
}

func ChangePassword(ctx *gin.Context) {
	var (
		userRepo = NewRepository(connection.DBConnections)
		userSrv  = NewService(userRepo)
	)

	err := userSrv.ChangePasswordService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "awesome, successfully change password")
}

func GetListUser(ctx *gin.Context) {
	var (
		userRepo = NewRepository(connection.DBConnections)
		userSrv  = NewService(userRepo)
	)

	users, err := userSrv.GetListUserService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get all user data", users)
}
