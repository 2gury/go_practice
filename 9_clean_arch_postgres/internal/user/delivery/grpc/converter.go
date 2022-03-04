package grpc

import (
	"go_practice/9_clean_arch_db/internal/models"
)

func GrpcUserToModel(user *User) *models.User {
	return &models.User{
		Id: user.Id,
		Email: user.Email,
		Password: user.Password,
		Role: user.Role,
	}
}

func ModelUserToGrpc(user *models.User) *User {
	return &User{
		Id: user.Id,
		Email: user.Email,
		Password: user.Password,
		Role: user.Role,
	}
}
