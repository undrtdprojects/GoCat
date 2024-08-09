package user

import (
	"GoCat/helpers/common"
	"GoCat/middlewares"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetListUserService(ctx *gin.Context) (result []User, err error)
	GetUserByUsernameService(ctx *gin.Context) (result User, err error)
	GetUserByIdService(ctx *gin.Context) (result User, err error)
	LoginService(ctx *gin.Context) (result LoginResponse, err error)
	SignUpService(ctx *gin.Context) (err error)
	UpdateUserService(ctx *gin.Context) (err error)
	ChangePasswordService(ctx *gin.Context) (err error)
	DeleteUserService(ctx *gin.Context) (err error)
}

type userService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &userService{repository}
}

func (service *userService) GetListUserService(ctx *gin.Context) (result []User, err error) {
	result, err = service.repository.GetList()

	if err != nil {
		return result, err
	}
	return result, nil
}

func (service *userService) GetUserByUsernameService(ctx *gin.Context) (result User, err error) {
	var userReq User
	err = ctx.ShouldBind(&userReq)
	if err != nil {
		return
	}

	result, err = service.repository.GetUserByUsername(userReq.Username)
	if err != nil {
		return
	}
	return
}

func (service *userService) GetUserByIdService(ctx *gin.Context) (result User, err error) {
	var userReq User
	err = ctx.ShouldBind(&userReq)
	if err != nil {
		return
	}

	result, err = service.repository.GetUserById(userReq.Id)
	if err != nil {
		return
	}
	return
}

func (service *userService) LoginService(ctx *gin.Context) (result LoginResponse, err error) {
	var userReq LoginRequest

	err = ctx.ShouldBind(&userReq)
	if err != nil {
		return
	}

	err = userReq.ValidateLogin()
	if err != nil {
		return
	}

	user, err := service.repository.Login(userReq)
	if err != nil {
		return
	}

	if common.IsEmptyField(user.Id) {
		err = errors.New("invalid account")
		return
	}

	matches := common.CheckPassword(user.Password, userReq.Password)
	if !matches {
		err = errors.New("wrong username or password")
		return
	}

	jwtToken, err := middlewares.GenerateJwtToken(user.Id, user.Username, user.RoleId)
	if err != nil {
		return
	}

	result.Token = jwtToken

	return
}

func (service *userService) SignUpService(ctx *gin.Context) (err error) {
	var userReq SignUpRequest

	err = ctx.ShouldBind(&userReq)
	if err != nil {
		return err
	}

	err = userReq.ValidateSignUp()
	if err != nil {
		return err
	}

	user, err := userReq.ConvertToModelForSignUp(ctx)
	if err != nil {
		return err
	}

	fmt.Println("user:", user)

	err = service.repository.SignUp(user)
	if err != nil {
		return err
	}

	return nil
}

func (service *userService) UpdateUserService(ctx *gin.Context) (err error) {
	var userReq User
	err = ctx.ShouldBind(&userReq)
	if err != nil {
		return err
	}

	err = service.repository.Update(userReq)
	if err != nil {
		return err
	}
	return nil
}

func (service *userService) ChangePasswordService(ctx *gin.Context) (err error) {
	var userReq User
	err = ctx.ShouldBind(&userReq)
	if err != nil {
		return err
	}

	err = service.repository.ChangePassword(userReq)
	if err != nil {
		return err
	}
	return nil
}

func (service *userService) DeleteUserService(ctx *gin.Context) (err error) {
	var userReq User
	err = ctx.ShouldBind(&userReq)
	if err != nil {
		return err
	}
	err = service.repository.Delete(userReq)
	if err != nil {
		return err
	}
	return nil
}
