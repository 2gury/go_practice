package usecases

import (
	"context"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/user"
	"go_practice/9_clean_arch_db/internal/user/delivery/grpc"
	"go_practice/9_clean_arch_db/tools/password"
)

type UserUsecase struct {
	userSvc grpc.UserServiceClient
}

func NewUserUsecase(svc grpc.UserServiceClient) user.UserUsecase {
	return &UserUsecase{
		userSvc: svc,
	}
}

func (u *UserUsecase) GetById(id uint64) (*models.User, *errors.Error) {
	usr, err := u.userSvc.GetById(context.Background(), &grpc.UserIdValue{Value: id})
	if err != nil {
		return nil, errors.GetCustomError(err)
	}

	return grpc.GrpcUserToModel(usr), nil
}

func (u *UserUsecase) GetByEmail(email string) (*models.User, *errors.Error) {
	usr, err := u.userSvc.GetByEmail(context.Background(), &grpc.EmailValue{Value: email})
	if err != nil {
		return nil, errors.GetCustomError(err)
	}

	return grpc.GrpcUserToModel(usr), nil
}

func (u *UserUsecase) Create(usr *models.User) (uint64, *errors.Error) {
	lastId, err := u.userSvc.Create(context.Background(), grpc.ModelUserToGrpc(usr))
	if err != nil {
		return 0, errors.GetCustomError(err)
	}

	return lastId.Value, nil
}

func (u *UserUsecase) UpdateUserPassword(usr *models.User) *errors.Error {
	_, err := u.userSvc.UpdateUserPassword(context.Background(), grpc.ModelUserToGrpc(usr))
	if err != nil {
		return errors.GetCustomError(err)
	}
	return nil
}

func (u *UserUsecase) DeleteUserById(id uint64) *errors.Error {
	_, err := u.userSvc.DeleteUserById(context.Background(), &grpc.UserIdValue{Value: id})
	if err != nil {
		return errors.GetCustomError(err)
	}

	return nil
}

func (u *UserUsecase) ComparePasswords(password, repeatedPassword string) *errors.Error {
	ok := password == repeatedPassword
	if !ok {
		return errors.Get(consts.CodeUserPasswordsDoNotMatch)
	}

	return nil
}

func (u *UserUsecase) ComparePasswordAndHash(usr *models.User, pass string) *errors.Error {
	ok := password.VerifyPasswordAndHash(pass, usr.Password)
	if !ok {
		return errors.Get(consts.CodeWrongPasswords)
	}

	return nil
}
