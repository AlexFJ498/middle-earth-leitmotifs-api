package dto

import domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"

type CategoryCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type CategoryUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}

type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewCategoryResponse(category domain.Category) CategoryResponse {
	return CategoryResponse{
		ID:   category.ID().String(),
		Name: category.Name().String(),
	}
}
