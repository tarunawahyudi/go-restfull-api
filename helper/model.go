package helper

import (
	"belajar-golang-restfull-api/model/domain"
	"belajar-golang-restfull-api/model/dto"
)

func ToCategoryResponse(category domain.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []dto.CategoryResponse {
	var categoryResponses []dto.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}
