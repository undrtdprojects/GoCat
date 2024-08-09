package transaction0

import (
	"GoCat/helpers/common"
	"GoCat/middlewares"
	"GoCat/modules/menu"
	"GoCat/modules/payment"
	"GoCat/modules/transaction1"
	"GoCat/modules/user"
	"errors"
	"fmt"
	"time"

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
	menuRepo    menu.Repository
	trans1Repo  transaction1.Repository
}

func NewService(repository Repository, repoUser user.Repository, paymentRepo payment.Repository, menuRepo menu.Repository, trans1Repo transaction1.Repository) Service {
	return &transaction0Service{repository, repoUser, paymentRepo, menuRepo, trans1Repo}
}

func (service *transaction0Service) CreateTransaction0Service(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	var newTransaction0 Transaction0

	err = ctx.ShouldBind(&newTransaction0)
	if err != nil {
		return err
	}

	user, err := service.repoUser.GetUserByUsername(userLogin.Username)
	if err != nil {
		return err
	}
	newTransaction0.UserId = user.Id

	if common.IsEmptyField(user.Username) {
		return errors.New("user not registered")
	}

	payment, err := service.paymentRepo.GetPaymentByIdRepository(newTransaction0.PaymentId)
	if err != nil {
		return err
	}

	fmt.Println("payment: ", payment)

	if common.IsEmptyField(payment.Name) {
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

	var GrandTotalPrice int
	// insert di table transaction1
	for _, trans1 := range newTransaction0.ListDetail {
		trans1.TransactionId = newTransaction0.Id
		trans1.DateTransaction = time.Now()

		menu, err := service.menuRepo.GetMenuByIdRepository(trans1.MenuId)
		if err != nil {
			return err
		}
		price := menu.Price * trans1.Qty
		trans1.TotalPrice = price
		GrandTotalPrice += price

		trans1.CreatedAt = defaultField.CreatedAt
		trans1.CreatedBy = userLogin.Username
		trans1.CreatedOn = defaultField.CreatedOn
		trans1.ModifiedAt = defaultField.ModifiedAt
		trans1.ModifiedBy = userLogin.Username
		trans1.ModifiedOn = defaultField.ModifiedOn

		err = service.trans1Repo.CreateTransaction1Repository(trans1)
		if err != nil {
			return err
		}
	}
	newTransaction0.GrandTotalPrice = GrandTotalPrice

	fmt.Println("newTransaction0: ", newTransaction0)
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
