package transaction0

import (
	"GoCat/helpers/common"
	"GoCat/middlewares"
	"GoCat/modules/payment"
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
	repository  Repository
	repoUser    user.Repository
	paymentRepo payment.Repository
}

type userService struct {
	userRepository user.Repository
}

func NewService(repository Repository, repoUser user.Repository, paymentRepo payment.Repository) Service {
	return &transaction0Service{repository, repoUser, paymentRepo}
}

func (service *transaction0Service) CreateTransaction0Service(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

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

	payment, err := service.paymentRepo.GetPaymentByIdRepository(newTransaction0.PaymentId)
	if err != nil {
		return err
	}
	if common.IsEmptyField(payment.Id) {
		return errors.New("payment_id not registered")
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
}

func (service *transaction0Service) GetAllTransaction0Service(ctx *gin.Context) (transaction0s []Transaction0, err error) {
	return service.repository.GetAllTransaction0Repository()
}

func (service *transaction0Service) GetTransaction0ByIdService(ctx *gin.Context) (transaction0 Transaction0, err error) {
	var id = ctx.Param("id")

	return service.repository.GetTransaction0ByIdRepository(id)
}

func (service *transaction0Service) DeleteTransaction0Service(ctx *gin.Context) (err error) {
	var transaction0 Transaction0
	transaction0.Id = ctx.Param("id")

	return service.repository.DeleteTransaction0Repository(transaction0)
}

func (service *transaction0Service) UpdateTransaction0Service(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

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
}
