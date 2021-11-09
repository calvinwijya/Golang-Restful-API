package service

import (
	"belajar-golang-resful-api/helper"
	"belajar-golang-resful-api/model/domain"
	"belajar-golang-resful-api/model/domain/web"
	"belajar-golang-resful-api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRespository
	DB                 *sql.DB
	Validation         *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRespository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validation:         validate,
	}
}

func (service *CategoryServiceImpl) CreateCategory(ctx context.Context, r web.CategoryCreateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: r.Name,
	}

	service.CategoryRepository.Save(ctx, tx, category)

	return helper.CategoryToResponse(category)
}
func (service *CategoryServiceImpl) UpdateCategory(ctx context.Context, r web.CategoryUpdateRequest) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, r.Id)
	helper.PanicIfError(err)

	category.Name = r.Name

	service.CategoryRepository.Update(ctx, tx, category)

	return helper.CategoryToResponse(category)
}
func (service *CategoryServiceImpl) DeleteCategory(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.CategoryRepository.Delete(ctx, tx, category)

}
func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.CategoryToResponse(category)
}
func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
