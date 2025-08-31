package dto

import domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"

type GroupCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type GroupUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}

type GroupResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewGroupResponse(group domain.Group) GroupResponse {
	return GroupResponse{
		ID:   group.ID().String(),
		Name: group.Name().String(),
	}
}
