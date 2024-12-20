package service

import (
	"belajar-golang-restfull-api/model/dto"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, request dto.CategoryCreateRequest) dto.CategoryResponse
	Update(ctx context.Context, request dto.CategoryUpdateRequest) dto.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) dto.CategoryResponse
	FindAll(ctx context.Context) []dto.CategoryResponse
}
