package usecases

import (
	"database/sql"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/user"
	"go_practice/9_clean_arch_db/tools/password"
)

type UserUsecase struct {
	userRep user.UserRepository
}

func NewUserUsecase(rep user.UserRepository) user.UserUsecase {
	return &UserUsecase{
		userRep: rep,
	}
}

func (u *UserUsecase) GetById(id uint64) (*models.User, *errors.Error) {
	usr, err := u.userRep.SelectById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Get(consts.CodeUserDoesNotExist)
		}
		return nil, errors.Get(consts.CodeInternalError)
	}

	return usr, nil
}

func (u *UserUsecase) GetByEmail(email string) (*models.User, *errors.Error) {
	usr, err := u.userRep.SelectByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Get(consts.CodeUserDoesNotExist)
		}
		return nil, errors.Get(consts.CodeInternalError)
	}

	return usr, nil
}

func (u *UserUsecase) Create(usr *models.User) (uint64, *errors.Error) {
	hashedPassword, err := password.HashPassword(usr.Password)
	if err != nil {
		return 0, errors.Get(consts.CodeInternalError)
	}
	usr.Password = hashedPassword

	lastId, err := u.userRep.Insert(usr)
	if err != nil {
		return 0, errors.Get(consts.CodeInternalError)
	}

	return lastId, nil
}

func (u *UserUsecase) UpdateUserPassword(usr *models.User) *errors.Error {
	hashedPassword, err := password.HashPassword(usr.Password)
	if err != nil {
		return errors.Get(consts.CodeInternalError)
	}
	usr.Password = hashedPassword

	err = u.userRep.UpdatePassword(usr)
	if err != nil {
		return errors.Get(consts.CodeInternalError)
	}
	return nil

}

func (u *UserUsecase) DeleteUserById(id uint64) *errors.Error {
	err := u.userRep.DeleteById(id)
	if err != nil {
		return errors.Get(consts.CodeInternalError)
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
