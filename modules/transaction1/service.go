package transaction1

import (
	"GoCat/helpers/common"
	"GoCat/middlewares"
	"GoCat/modules/menu"

	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateTransaction1Service(ctx *gin.Context) (err error)
	GetAllTransaction1Service(ctx *gin.Context) (result []Transaction1, err error)
	GetTransaction1ByIdService(ctx *gin.Context) (result Transaction1, err error)
	DeleteTransaction1Service(ctx *gin.Context) (err error)
	UpdateTransaction1Service(ctx *gin.Context) (err error)
}

type transaction1Service struct {
	repository Repository
	menuRepo   menu.Repository
}

func NewService(repository Repository, menuRepo menu.Repository) Service {
	return &transaction1Service{repository, menuRepo}
}

func (service *transaction1Service) CreateTransaction1Service(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	var newTransaction1 Transaction1

	err = ctx.ShouldBind(&newTransaction1)
	if err != nil {
		return err
	}

	menu, err := service.menuRepo.GetMenuByIdRepository(newTransaction1.MenuId)
	if err != nil {
		return err
	}
	if common.IsEmptyField(menu.Id) {
		return errors.New("menu_id not registered")
	}

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField(ctx)

	newTransaction1.CreatedAt = defaultField.CreatedAt
	newTransaction1.CreatedBy = userLogin.Username
	newTransaction1.CreatedOn = defaultField.CreatedOn
	newTransaction1.ModifiedAt = defaultField.ModifiedAt
	newTransaction1.ModifiedBy = userLogin.Username
	newTransaction1.ModifiedOn = defaultField.ModifiedOn

	err = service.repository.CreateTransaction1Repository(newTransaction1)
	if err != nil {
		return errors.New("failed to add new transaction1")
	}

	return nil
}

func (service *transaction1Service) GetAllTransaction1Service(ctx *gin.Context) (transaction1s []Transaction1, err error) {
	return service.repository.GetAllTransaction1Repository()
}

func (service *transaction1Service) GetTransaction1ByIdService(ctx *gin.Context) (transaction1 Transaction1, err error) {
	var (
		idInt int
		id    = ctx.Param("id")
	)

	idInt, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id role from param")
		return
	}

	return service.repository.GetTransaction1ByIdRepository(idInt)
}

func (service *transaction1Service) DeleteTransaction1Service(ctx *gin.Context) (err error) {
	var (
		transaction1 Transaction1
		id           = ctx.Param("id")
	)

	transaction1.Id, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id transaction1 from param")
		return
	}

	return service.repository.DeleteTransaction1Repository(transaction1)
}

func (service *transaction1Service) UpdateTransaction1Service(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	var (
		newTransaction1 Transaction1
		id              = ctx.Param("id")
	)

	err = ctx.ShouldBind(&newTransaction1)
	if err != nil {
		return
	}

	newTransaction1.Id, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id transaction1 from param")
		return
	}

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField(ctx)

	newTransaction1.ModifiedAt = defaultField.ModifiedAt
	newTransaction1.ModifiedBy = userLogin.Username
	newTransaction1.ModifiedOn = defaultField.ModifiedOn

	return service.repository.UpdateTransaction1Repository(newTransaction1)
}
