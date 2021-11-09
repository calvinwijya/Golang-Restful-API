package service

import (
	"belajar-golang-resful-api/model/domain/web"
	"context"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, r web.CategoryCreateRequest) web.CategoryResponse
	UpdateCategory(ctx context.Context, r web.CategoryUpdateRequest) web.CategoryResponse
	DeleteCategory(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
