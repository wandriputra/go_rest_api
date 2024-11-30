package services

import (
	"context"
	"encoding/json"
	"os"

	"github.com/wandriputra/go_rest_api/src/domains"
)

type userService struct{}

func NewUserService() domains.IUserService {
	return &userService{}
}

func (service userService) GetAllUsers(ctx context.Context) (response any, err error) {
	file, err := os.Open("public/input.json")

	if err != nil {
		return response, err
	}

	defer file.Close()

	var users interface{}
	err = json.NewDecoder(file).Decode(&users)

	if err != nil {
		return nil, err // Return the error if JSON parsing fails
	}
	return users, nil
}

func (service userService) GetSingleUser(ctx context.Context, id domains.GetSingleUserRequest) (response any, err error) {
	return response, nil
}

func (service userService) CreateUser(ctx context.Context, createUser domains.CreateUserRequest) (response any, err error) {
	return response, nil
}
