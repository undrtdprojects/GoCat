package categories

import (
	"errors"
	"fmt"
	"quiz-3-sanbercode-greg/helpers/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateCategoriesService(ctx *gin.Context) (err error)
	GetAllCategoriesService(ctx *gin.Context) (result []Categories, err error)
	GetCategoriesByIdService(ctx *gin.Context) (result Categories, err error)
	DeleteCategoriesService(ctx *gin.Context) (err error)
	UpdateCategoriesService(ctx *gin.Context) (err error)
}

type categoryService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &categoryService{repository}
}

func (service *categoryService) CreateCategoriesService(ctx *gin.Context) (err error) {
	var newCategories Categories

	err = ctx.ShouldBind(&newCategories)
	if err != nil {
		return err
	}

	var categories []Categories
	categories, err = service.repository.GetAllCategoriesRepository()
	if err != nil {
		return err
	}

	category, err := service.repository.GetCategoriesByNameRepository(newCategories.Name)
	if err != nil {
		return err
	}

	if len(categories) != 0 && category.Name != "" {
		err = errors.New("category already exists")
		return err
	}

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField()

	newCategories.CreatedAt = defaultField.CreatedAt
	newCategories.CreatedBy = defaultField.CreatedBy
	newCategories.ModifiedAt = defaultField.ModifiedAt
	newCategories.ModifiedBy = defaultField.ModifiedBy

	fmt.Println("create categories :", newCategories)
	err = service.repository.CreateCategoriesRepository(newCategories)
	if err != nil {
		return errors.New("failed to add new category")
	}

	return
}

func (service *categoryService) GetAllCategoriesService(ctx *gin.Context) (categories []Categories, err error) {
	return service.repository.GetAllCategoriesRepository()
}

func (service *categoryService) GetCategoriesByIdService(ctx *gin.Context) (category Categories, err error) {
	var (
		idInt int
		id    = ctx.Param("id")
	)

	idInt, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id category from param")
		return
	}

	return service.repository.GetCategoriesByIdRepository(idInt)
}

func (service *categoryService) DeleteCategoriesService(ctx *gin.Context) (err error) {
	var (
		category Categories
		id       = ctx.Param("id")
	)

	category.Id, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id category from param")
		return
	}

	return service.repository.DeleteCategoriesRepository(category)
}

func (service *categoryService) UpdateCategoriesService(ctx *gin.Context) (err error) {
	var (
		category Categories
		id       = ctx.Param("id")
	)

	err = ctx.ShouldBind(&category)
	if err != nil {
		return
	}

	category.Id, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id category from param")
		return
	}
	return service.repository.UpdateCategoriesRepository(category)
}
