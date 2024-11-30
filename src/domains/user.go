package domains

import "context"

type IUserService interface {
	GetAllUsers(ctx context.Context) (users any, erro error)
	GetSingleUser(ctx context.Context, id GetSingleUserRequest) (user any, erro error)
	CreateUser(ctx context.Context, createUser CreateUserRequest) (users any, erro error)
}

type GetAllUsersResponse struct {
	Users any `json:"users"`
}

type GetSingleUserRequest struct {
	ID string `json:"id"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
