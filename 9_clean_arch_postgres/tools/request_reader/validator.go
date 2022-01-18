package request_reader

import (
	"github.com/asaskevich/govalidator"
	"go_practice/8_clean_arch/internal/consts"
	"go_practice/8_clean_arch/internal/helpers/errors"
)

func ValidateStruct(request interface{}) *errors.Error {
	ok, err := govalidator.ValidateStruct(request)
	if err != nil {
		return errors.New(consts.CodeValidateError, err)
	}
	if !ok {
		return errors.Get(consts.CodeValidateError)
	}
	return nil
}
