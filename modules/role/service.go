package role

import (
	"GoCat/helpers/common"
	"GoCat/middlewares"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateRoleService(ctx *gin.Context) (err error)
	GetAllRoleService(ctx *gin.Context) (result []Role, err error)
	GetRoleByIdService(ctx *gin.Context) (result Role, err error)
	DeleteRoleService(ctx *gin.Context) (err error)
	UpdateRoleService(ctx *gin.Context) (err error)
}

type roleService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &roleService{repository}
}

func (service *roleService) CreateRoleService(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	var newRole Role

	err = ctx.ShouldBind(&newRole)
	if err != nil {
		return err
	}

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField(ctx)

	newRole.CreatedAt = defaultField.CreatedAt
	newRole.CreatedBy = userLogin.Username
	newRole.CreatedOn = defaultField.CreatedOn
	newRole.ModifiedAt = defaultField.ModifiedAt
	newRole.ModifiedBy = userLogin.Username
	newRole.ModifiedOn = defaultField.ModifiedOn

	err = service.repository.CreateRoleRepository(newRole)
	if err != nil {
		return errors.New("failed to add new role")
	}

	return nil
}

func (service *roleService) GetAllRoleService(ctx *gin.Context) (roles []Role, err error) {
	return service.repository.GetAllRoleRepository()
}

func (service *roleService) GetRoleByIdService(ctx *gin.Context) (role Role, err error) {
	var (
		idInt int
		id    = ctx.Param("id")
	)

	idInt, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id role from param")
		return
	}

	return service.repository.GetRoleByIdRepository(idInt)
}

func (service *roleService) DeleteRoleService(ctx *gin.Context) (err error) {
	var (
		role Role
		id   = ctx.Param("id")
	)

	role.Id, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id role from param")
		return
	}

	return service.repository.DeleteRoleRepository(role)
}

func (service *roleService) UpdateRoleService(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	var (
		newRole Role
		id      = ctx.Param("id")
	)

	err = ctx.ShouldBind(&newRole)
	if err != nil {
		return
	}

	newRole.Id, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id role from param")
		return
	}

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField(ctx)

	newRole.ModifiedAt = defaultField.ModifiedAt
	newRole.ModifiedBy = userLogin.Username
	newRole.ModifiedOn = defaultField.ModifiedOn

	return service.repository.UpdateRoleRepository(newRole)
}
