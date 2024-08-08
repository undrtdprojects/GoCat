package transaction0

import (
	"GoCat/helpers/common"
	"GoCat/helpers/constant"
	"GoCat/middlewares"
	"GoCat/modules/user"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateTransaction0Service(ctx *gin.Context) (err error)
	GetAllTransaction0Service(ctx *gin.Context) (result []Transaction0, err error)
	GetTransaction0ByIdService(ctx *gin.Context) (result Transaction0, err error)
	DeleteTransaction0Service(ctx *gin.Context) (err error)
	UpdateTransaction0Service(ctx *gin.Context) (err error)
}

type transaction0Service struct {
	repository Repository
	repoUser   user.Repository
}

type userService struct {
	userRepository user.Repository
}

func NewService(repository Repository, repoUser user.Repository) Service {
	return &transaction0Service{repository, repoUser}
}

func (service *transaction0Service) CreateTransaction0Service(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	if common.CheckRole(userLogin.RoleId, constant.ReadActionUser.String()) {
		var newTransaction0 Transaction0

		err = ctx.ShouldBind(&newTransaction0)
		if err != nil {
			return err
		}

		user, err := service.repoUser.GetUserById(newTransaction0.UserId)
		if err != nil {
			return err
		}

		if common.IsEmptyField(user.Id) {
			return errors.New("user not registered")
		}

		defaultField := common.DefaultFieldTable{}
		defaultField.SetDefaultField(ctx)

		newTransaction0.CreatedAt = defaultField.CreatedAt
		newTransaction0.CreatedBy = userLogin.Username
		newTransaction0.CreatedOn = defaultField.CreatedOn
		newTransaction0.ModifiedAt = defaultField.ModifiedAt
		newTransaction0.ModifiedBy = userLogin.Username
		newTransaction0.ModifiedOn = defaultField.ModifiedOn

		index, err := service.repository.GetTransaction0CountRepository()
		if err != nil {
			return err
		}

		newTransaction0.Id = fmt.Sprintf("%s-%05d", "CAT", index)

		err = service.repository.CreateTransaction0Repository(newTransaction0)
		if err != nil {
			return errors.New("failed to add new transaction0")
		}

		return nil
	} else {
		return errors.New("you are not authorized")
	}
}

func (service *transaction0Service) GetAllTransaction0Service(ctx *gin.Context) (transaction0s []Transaction0, err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	if common.CheckRole(userLogin.RoleId, constant.ReadActionUser.String()) {
		return service.repository.GetAllTransaction0Repository()
	} else {
		return nil, errors.New("you are not authorized")
	}
}

func (service *transaction0Service) GetTransaction0ByIdService(ctx *gin.Context) (transaction0 Transaction0, err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	if common.CheckRole(userLogin.RoleId, constant.ReadActionUser.String()) {
		var id = ctx.Param("id")

		return service.repository.GetTransaction0ByIdRepository(id)
	} else {
		return transaction0, errors.New("you are not authorized")
	}
}

func (service *transaction0Service) DeleteTransaction0Service(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	if common.CheckRole(userLogin.RoleId, constant.DeleteActionUser.String()) {
		var transaction0 Transaction0

		transaction0.Id = ctx.Param("id")

		return service.repository.DeleteTransaction0Repository(transaction0)
	} else {
		return errors.New("you are not authorized")
	}
}

func (service *transaction0Service) UpdateTransaction0Service(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	if common.CheckRole(userLogin.RoleId, constant.UpdateActionUser.String()) {
		var newTransaction0 Transaction0

		err = ctx.ShouldBind(&newTransaction0)
		if err != nil {
			return
		}

		defaultField := common.DefaultFieldTable{}
		defaultField.SetDefaultField(ctx)

		newTransaction0.ModifiedAt = defaultField.ModifiedAt
		newTransaction0.ModifiedBy = userLogin.Username
		newTransaction0.ModifiedOn = defaultField.ModifiedOn

		newTransaction0.Id = ctx.Param("id")

		return service.repository.UpdateTransaction0Repository(newTransaction0)
	} else {
		return errors.New("you are not authorized")
	}
}
