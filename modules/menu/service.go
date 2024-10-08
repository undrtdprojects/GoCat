package menu

import (
	"GoCat/helpers/common"
	"GoCat/middlewares"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateMenuService(ctx *gin.Context) (err error)
	GetAllMenuService(ctx *gin.Context) (result []Menu, err error)
	GetMenuByIdService(ctx *gin.Context) (result Menu, err error)
	DeleteMenuService(ctx *gin.Context) (err error)
	UpdateMenuService(ctx *gin.Context) (err error)
}

type MenuService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &MenuService{repository}
}

func (service *MenuService) CreateMenuService(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	var newMenu Menu

	err = ctx.ShouldBind(&newMenu)
	if err != nil {
		return err
	}

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField(ctx)

	newMenu.CreatedAt = defaultField.CreatedAt
	newMenu.CreatedBy = userLogin.Username
	newMenu.CreatedOn = defaultField.CreatedOn
	newMenu.ModifiedAt = defaultField.ModifiedAt
	newMenu.ModifiedBy = userLogin.Username
	newMenu.ModifiedOn = defaultField.ModifiedOn

	index, err := service.repository.GetMenuCountByCategoryIdRepository(newMenu.CategoryId)
	if err != nil {
		return err
	}

	if common.IsEmptyField(index) {
		return errors.New("category not found")
	}

	newMenu.Id = fmt.Sprintf("%s-%04d", newMenu.CategoryId, index)

	err = service.repository.CreateMenuRepository(newMenu)
	if err != nil {
		return errors.New("failed to add new Menu")
	}

	return nil
}

func (service *MenuService) GetAllMenuService(ctx *gin.Context) (categories []Menu, err error) {
	return service.repository.GetAllMenuRepository()
}

func (service *MenuService) GetMenuByIdService(ctx *gin.Context) (menu Menu, err error) {
	var Id = ctx.Param("id")

	menu, err = service.repository.GetMenuByIdRepository(Id)
	if common.IsEmptyField(menu.Id) {
		return menu, errors.New("menu not found")
	}

	return menu, err
}

func (service *MenuService) DeleteMenuService(ctx *gin.Context) (err error) {
	var newMenu Menu
	newMenu.Id = ctx.Param("id")

	return service.repository.DeleteMenuRepository(newMenu)
}

func (service *MenuService) UpdateMenuService(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	var newMenu Menu

	err = ctx.ShouldBind(&newMenu)
	if err != nil {
		return
	}

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField(ctx)

	newMenu.ModifiedAt = defaultField.ModifiedAt
	newMenu.ModifiedBy = userLogin.Username
	newMenu.ModifiedOn = defaultField.ModifiedOn

	newMenu.Id = ctx.Param("id")
	return service.repository.UpdateMenuRepository(newMenu)
}
