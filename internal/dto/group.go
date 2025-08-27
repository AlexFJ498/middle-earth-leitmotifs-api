package dto

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

func NewGroupResponse(id, name string) GroupResponse {
	return GroupResponse{
		ID:   id,
		Name: name,
	}
}
