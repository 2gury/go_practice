package usecases

import "go_practice/9_clean_arch_db/internal/user"

type UserUsecase struct {
	userRep user.UserRepository
}

func NewUserUsecase(rep user.UserRepository) user.UserUsecase {
	return UserUsecase{
		userRep: rep,
	}
}
