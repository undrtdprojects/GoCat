package menu

import (
	"errors"
	"fmt"
	"quiz-3-sanbercode-greg/helpers/common"
	"strconv"

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
	var newMenu Menu

	err = ctx.ShouldBind(&newMenu)
	if err != nil {
		return err
	}

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField()

	newMenu.CreatedAt = defaultField.CreatedAt
	newMenu.CreatedBy = defaultField.CreatedBy
	newMenu.ModifiedAt = defaultField.ModifiedAt
	newMenu.ModifiedBy = defaultField.ModifiedBy

	fmt.Println("create categories :", newMenu)
	err = service.repository.CreateMenuRepository(newMenu)
	if err != nil {
		return errors.New("failed to add new Menu")
	}

	return
}

func (service *MenuService) GetAllMenuService(ctx *gin.Context) (categories []Menu, err error) {
	return service.repository.GetAllMenuRepository()
}

func (service *MenuService) GetMenuByIdService(ctx *gin.Context) (Menu Menu, err error) {
	var (
		idInt int
		id    = ctx.Param("id")
	)

	idInt, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id Menu from param")
		return
	}

	return service.repository.GetMenuByIdRepository(idInt)
}

func (service *MenuService) GetAllMenusByMenuService(ctx *gin.Context) (Menus []Menu.Menu, err error) {
	var (
		idInt int
		id    = ctx.Param("id")
		name  = ctx.Param("name")
	)

	idInt, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id Menu from param")
		return
	}

	return service.repository.GetAllMenusByMenuRepository(idInt, name)
}

func (service *MenuService) DeleteMenuService(ctx *gin.Context) (err error) {
	var (
		Menu Menu
		id   = ctx.Param("id")
	)

	Menu.Id, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id Menu from param")
		return
	}

	return service.repository.DeleteMenuRepository(Menu)
}

func (service *MenuService) UpdateMenuService(ctx *gin.Context) (err error) {
	var (
		Menu Menu
		id   = ctx.Param("id")
	)

	err = ctx.ShouldBind(&Menu)
	if err != nil {
		return
	}

	Menu.Id, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id Menu from param")
		return
	}
	return service.repository.UpdateMenuRepository(Menu)
}
