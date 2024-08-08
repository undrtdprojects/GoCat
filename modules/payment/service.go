package payment

import (
	"GoCat/helpers/common"
	"GoCat/helpers/constant"
	"GoCat/middlewares"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreatePaymentService(ctx *gin.Context) (err error)
	GetAllPaymentService(ctx *gin.Context) (result []Payment, err error)
	GetPaymentByIdService(ctx *gin.Context) (result Payment, err error)
	DeletePaymentService(ctx *gin.Context) (err error)
	UpdatePaymentService(ctx *gin.Context) (err error)
}

type paymentService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &paymentService{repository}
}

func (service *paymentService) CreatePaymentService(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	if common.CheckRole(userLogin.RoleId, constant.ReadActionUser.String()) {
		var newPayment Payment

		err = ctx.ShouldBind(&newPayment)
		if err != nil {
			return err
		}

		defaultField := common.DefaultFieldTable{}
		defaultField.SetDefaultField(ctx)

		newPayment.CreatedAt = defaultField.CreatedAt
		newPayment.CreatedBy = userLogin.Username
		newPayment.CreatedOn = defaultField.CreatedOn
		newPayment.ModifiedAt = defaultField.ModifiedAt
		newPayment.ModifiedBy = userLogin.Username
		newPayment.ModifiedOn = defaultField.ModifiedOn

		err = service.repository.CreatePaymentRepository(newPayment)
		if err != nil {
			return errors.New("failed to add new payment")
		}

		return nil
	} else {
		return errors.New("you are not authorized")
	}
}

func (service *paymentService) GetAllPaymentService(ctx *gin.Context) (payments []Payment, err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	if common.CheckRole(userLogin.RoleId, constant.ReadActionUser.String()) {
		return service.repository.GetAllPaymentRepository()
	} else {
		return nil, errors.New("you are not authorized")
	}
}

func (service *paymentService) GetPaymentByIdService(ctx *gin.Context) (payment Payment, err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	if common.CheckRole(userLogin.RoleId, constant.ReadActionUser.String()) {
		var (
			idInt int
			id    = ctx.Param("id")
		)

		idInt, err = strconv.Atoi(id)
		if err != nil {
			err = errors.New("failed to get id payment from param")
			return
		}

		return service.repository.GetPaymentByIdRepository(idInt)
	} else {
		return payment, errors.New("you are not authorized")
	}
}

func (service *paymentService) DeletePaymentService(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	if common.CheckRole(userLogin.RoleId, constant.DeleteActionUser.String()) {
		var (
			payment Payment
			id      = ctx.Param("id")
		)

		payment.Id, err = strconv.Atoi(id)
		if err != nil {
			err = errors.New("failed to get id payment from param")
			return
		}

		return service.repository.DeletePaymentRepository(payment)
	} else {
		return errors.New("you are not authorized")
	}
}

func (service *paymentService) UpdatePaymentService(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	if common.CheckRole(userLogin.RoleId, constant.UpdateActionUser.String()) {
		var (
			newPayment Payment
			id         = ctx.Param("id")
		)

		err = ctx.ShouldBind(&newPayment)
		if err != nil {
			return
		}

		newPayment.Id, err = strconv.Atoi(id)
		if err != nil {
			err = errors.New("failed to get id payment from param")
			return
		}

		defaultField := common.DefaultFieldTable{}
		defaultField.SetDefaultField(ctx)

		newPayment.ModifiedAt = defaultField.ModifiedAt
		newPayment.ModifiedBy = userLogin.Username
		newPayment.ModifiedOn = defaultField.ModifiedOn

		return service.repository.UpdatePaymentRepository(newPayment)
	} else {
		return errors.New("you are not authorized")
	}
}
