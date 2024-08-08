package role

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
		api.POST("/role", CreateRoleRouter)
		api.GET("/role", GetAllRoleRouter)
		api.GET("/role/:id", GetRoleByIdRouter)
		api.PUT("/role/:id", UpdateRoleRouter)
		api.DELETE("/role/:id", DeleteRoleRouter)
	}
}

func CreateRoleRouter(ctx *gin.Context) {
	var (
		roleRepo = NewRepository(connection.DBConnections)
		roleSrv  = NewService(roleRepo)
	)

	err := roleSrv.CreateRoleService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully added role data")
}

func GetAllRoleRouter(ctx *gin.Context) {
	var (
		roleRepo = NewRepository(connection.DBConnections)
		roleSrv  = NewService(roleRepo)
	)

	roles, err := roleSrv.GetAllRoleService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get all role data", roles)
}

func GetRoleByIdRouter(ctx *gin.Context) {
	var (
		roleRepo = NewRepository(connection.DBConnections)
		roleSrv  = NewService(roleRepo)
	)

	role, err := roleSrv.GetRoleByIdService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get role data", role)
}

func DeleteRoleRouter(ctx *gin.Context) {
	var (
		roleRepo = NewRepository(connection.DBConnections)
		roleSrv  = NewService(roleRepo)
	)

	err := roleSrv.DeleteRoleService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully delete role data")
}

func UpdateRoleRouter(ctx *gin.Context) {
	var (
		roleRepo = NewRepository(connection.DBConnections)
		roleSrv  = NewService(roleRepo)
	)

	err := roleSrv.UpdateRoleService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully update role data")
}
