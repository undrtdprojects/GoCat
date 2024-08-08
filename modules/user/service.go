package user

import (
	"GoCat/helpers/common"
	"GoCat/helpers/constant"
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
	DeleteUserService(ctx *gin.Context) (err error)
}

type userService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &userService{
		repository,
	}
}

func (service *userService) GetListUserService(ctx *gin.Context) (result []User, err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)
	if common.CheckRole(userLogin.RoleId, constant.ReadActionUser.String()) {
		result, err = service.repository.GetList()
		if err != nil {
			return
		}
		return
	} else {
		return nil, errors.New("you don't have access to get all user")
	}
}

func (service *userService) GetUserByUsernameService(ctx *gin.Context) (result User, err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)
	if common.CheckRole(userLogin.RoleId, constant.ReadActionUser.String()) {
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
	} else {
		return result, errors.New("you don't have access to get user by username")
	}
}
func (service *userService) GetUserByIdService(ctx *gin.Context) (result User, err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	if common.CheckRole(userLogin.RoleId, constant.ReadActionUser.String()) {
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
	} else {
		return result, errors.New("you don't have access to get user by id")
	}
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
	fmt.Println("user :", user)
	err = service.repository.SignUp(user)
	if err != nil {
		return err
	}

	return nil
}

func (service *userService) UpdateUserService(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)
	if common.CheckRole(userLogin.RoleId, constant.UpdateActionUser.String()) {
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
	} else {
		return errors.New("you don't have access to update user")
	}
}

func (service *userService) DeleteUserService(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)
	if common.CheckRole(userLogin.RoleId, constant.DeleteActionUser.String()) {
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
	} else {
		return errors.New("you don't have access to delete user")
	}
}
