package grpc

import (
	"context"
	"database/sql"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/user"
	"go_practice/9_clean_arch_db/tools/password"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	userRep user.UserRepository
	UnimplementedUserServiceServer
}

func NewUserService(rep user.UserRepository) *UserService {
	return &UserService{
		userRep: rep,
	}
}

func (um *UserService) GetById(ctx context.Context, value *UserIdValue) (*User, error) {
	usr, err := um.userRep.SelectById(value.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.GetErrorFromGrpc(consts.CodeUserDoesNotExist, err)
		}
		return nil, errors.GetErrorFromGrpc(consts.CodeInternalError, err)
	}

	return ModelUserToGrpc(usr), nil
}

func (um *UserService) Create(ctx context.Context, user *User) (*UserIdValue, error) {
	hashedPassword, err := password.HashPassword(user.Password)
	if err != nil {
		return nil, errors.GetErrorFromGrpc(consts.CodeInternalError, err)
	}
	user.Password = hashedPassword

	lastId, err := um.userRep.Insert(GrpcUserToModel(user))
	if err != nil {
		return nil, errors.GetErrorFromGrpc(consts.CodeInternalError, err)
	}

	return &UserIdValue{Value: lastId}, nil
}

func (um *UserService) GetByEmail(ctx context.Context, email *EmailValue) (*User, error) {
	usr, err := um.userRep.SelectByEmail(email.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.GetErrorFromGrpc(consts.CodeUserDoesNotExist, err)
		}
		return nil, errors.GetErrorFromGrpc(consts.CodeInternalError, err)
	}

	return ModelUserToGrpc(usr), nil
}

func (um *UserService) UpdateUserPassword(ctx context.Context, user *User) (*emptypb.Empty, error) {
	hashedPassword, err := password.HashPassword(user.Password)
	if err != nil {
		return &emptypb.Empty{}, errors.GetErrorFromGrpc(consts.CodeInternalError, err)
	}
	user.Password = hashedPassword

	err = um.userRep.UpdatePassword(GrpcUserToModel(user))
	if err != nil {
		return &emptypb.Empty{}, errors.GetErrorFromGrpc(consts.CodeInternalError, err)
	}

	return &emptypb.Empty{}, nil
}

func (um *UserService) DeleteUserById(ctx context.Context, id *UserIdValue) (*emptypb.Empty, error) {
	err := um.userRep.DeleteById(id.Value)
	if err != nil {
		return &emptypb.Empty{}, errors.GetErrorFromGrpc(consts.CodeInternalError, err)
	}

	return &emptypb.Empty{}, nil
}

