package user

import (
	"GoCat/helpers/common"
	"errors"
	"os"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	RoleId     int       `json:"role_id"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	CreatedOn  string    `json:"created_on"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
	ModifiedOn string    `json:"modified_on"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleId   int    `json:"role_id"`
}

func (l *LoginRequest) ValidateLogin() (err error) {
	if common.IsEmptyField(l.Username) {
		return errors.New("username is required")
	}

	if common.IsEmptyField(l.Password) {
		return errors.New("password is required")
	}

	return
}

type LoginResponse struct {
	Token string `json:"token"`
}

type SignUpRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	ReTypePassword string `json:"re_type_password"`
	RoleId         int    `json:"role_id"`
}

func (s *SignUpRequest) ValidateSignUp() (err error) {
	if common.IsEmptyField(s.Username) {
		return errors.New("username is required")
	}

	if common.IsEmptyField(s.Password) {
		return errors.New("password is required")
	}
	if common.IsEmptyField(s.ReTypePassword) {
		return errors.New("retype password is required")
	}
	if s.Password != s.ReTypePassword {
		return errors.New("password and retype password not match")
	}
	if common.IsEmptyField(s.RoleId) {
		return errors.New("role is required")
	}

	re := regexp.MustCompile(`^(.{8,})$`)
	if !re.MatchString(s.Password) {
		return errors.New("please make sure that the password contains at least 8 character")
	}

	return nil
}

func (s *SignUpRequest) ConvertToModelForSignUp(ctx *gin.Context) (user User, err error) {
	hashedPassword, err := common.HashPassword(s.Password)
	if err != nil {
		err = errors.New("hashing password failed")
		return
	}

	hostname, _ := os.Hostname()

	return User{
		Username:   s.Username,
		Password:   hashedPassword,
		RoleId:     s.RoleId,
		CreatedAt:  time.Now(),
		CreatedOn:  hostname,
		ModifiedAt: time.Now(),
		ModifiedOn: hostname,
	}, nil
}

type SignUpResponse struct {
}
