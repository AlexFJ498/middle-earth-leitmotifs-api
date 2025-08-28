package dto

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

func NewCategoryResponse(id, name string) CategoryResponse {
	return CategoryResponse{
		ID:   id,
		Name: name,
	}
}
