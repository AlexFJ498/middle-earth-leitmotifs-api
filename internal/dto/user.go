package dto

import domain "github.com/AlexFJ498/middle-earth-leitmotifs-api/internal"

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserResponse(user domain.User) UserResponse {
	return UserResponse{
		ID:    user.ID().String(),
		Name:  user.Name().String(),
		Email: user.Email().String(),
	}
}
