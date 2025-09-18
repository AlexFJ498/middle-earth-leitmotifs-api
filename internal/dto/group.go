package dto

import domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"

type GroupCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageURL    string `json:"image_url" binding:"required"`
}

type GroupUpdateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageURL    string `json:"image_url" binding:"required"`
}

type GroupResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

func NewGroupResponse(group domain.Group) GroupResponse {
	return GroupResponse{
		ID:          group.ID().String(),
		Name:        group.Name().String(),
		Description: group.Description().String(),
		ImageURL:    group.ImageURL().String(),
	}
}
