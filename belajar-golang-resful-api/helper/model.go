package helper

import (
	"belajar-golang-resful-api/model/domain"
	"belajar-golang-resful-api/model/domain/web"
)

func CategoryToResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, CategoryToResponse(category))
	}
	return categoryResponses
}
