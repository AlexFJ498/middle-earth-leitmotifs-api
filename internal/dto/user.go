package dto

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

func NewUserResponse(id, name, email string) UserResponse {
	return UserResponse{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
