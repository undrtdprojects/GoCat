package categories

import (
	"GoCat/helpers/common"
	"GoCat/middlewares"
	"errors"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateCategoriesService(ctx *gin.Context) (err error)
	GetAllCategoriesService(ctx *gin.Context) (result []Categories, err error)
	GetCategoriesByIdService(ctx *gin.Context) (result Categories, err error)
	DeleteCategoriesService(ctx *gin.Context) (err error)
	UpdateCategoriesService(ctx *gin.Context) (err error)
}

type categoriesService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &categoriesService{repository}
}

func (service *categoriesService) CreateCategoriesService(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	var newCategories Categories

	err = ctx.ShouldBind(&newCategories)
	if err != nil {
		return err
	}

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField(ctx)

	newCategories.CreatedAt = defaultField.CreatedAt
	newCategories.CreatedBy = defaultField.CreatedOn
	newCategories.CreatedOn = userLogin.Username
	newCategories.ModifiedAt = defaultField.ModifiedAt
	newCategories.ModifiedBy = defaultField.ModifiedOn
	newCategories.ModifiedOn = userLogin.Username

	err = service.repository.CreateCategoriesRepository(newCategories)
	if err != nil {
		return errors.New(err.Error())
	}

	return
}

func (service *categoriesService) GetAllCategoriesService(ctx *gin.Context) (categories []Categories, err error) {
	return service.repository.GetAllCategoriesRepository()
}

func (service *categoriesService) GetCategoriesByIdService(ctx *gin.Context) (categories Categories, err error) {
	var id = ctx.Param("id")

	categories, err = service.repository.GetCategoriesByIdRepository(id)
	if common.IsEmptyField(categories.Id) {
		return categories, errors.New("categories not found")
	}

	return categories, err
}

func (service *categoriesService) DeleteCategoriesService(ctx *gin.Context) (err error) {
	var categories Categories
	categories.Id = ctx.Param("id")

	return service.repository.DeleteCategoriesRepository(categories)
}

func (service *categoriesService) UpdateCategoriesService(ctx *gin.Context) (err error) {
	userCtx, _ := ctx.Get("user")
	userLogin := userCtx.(*middlewares.Claims)

	var newCategories Categories

	err = ctx.ShouldBind(&newCategories)
	if err != nil {
		return
	}

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField(ctx)

	newCategories.ModifiedAt = defaultField.ModifiedAt
	newCategories.ModifiedBy = userLogin.Username
	newCategories.ModifiedOn = defaultField.ModifiedOn

	newCategories.Id = ctx.Param("id")
	return service.repository.UpdateCategoriesRepository(newCategories)
}
