package errors

import (
	"go_practice/8_clean_arch/internal/consts"
	"net/http"
)

type Error struct {
	HttpCode int    `json:"-"`
	Message  string `json:"message"`
	UserMessage string `json:"user_message"`
}

var WrongErrorCode = &Error{
	HttpCode: http.StatusTeapot,
	Message: "This error code doesn't exist",
	UserMessage: "Технические неполадки. Уже чиним",
}

func New(code consts.ErrorCode, err error) *Error {
	customErr, has := Errors[code]
	if !has {
		return WrongErrorCode
	}
	customErr.Message = err.Error()
	return customErr
}

func Get(code consts.ErrorCode) *Error {
	customErr, has := Errors[code]
	if !has {
		return WrongErrorCode
	}
	return customErr
}

var Errors =  map[consts.ErrorCode]*Error{
	consts.CodeBadRequest: {
		HttpCode: http.StatusBadRequest,
		Message: "This request format is invalid",
		UserMessage: "Неверный формат запроса",
	},
	consts.CodeInternalError: {
		HttpCode: http.StatusBadRequest,
		Message: "Sorry, can't handle request",
		UserMessage: "Что-то пошло не так",
	},
	consts.CodeProductDoesNotExist: {
		HttpCode: http.StatusBadRequest,
		Message: "Product doesn't exist",
		UserMessage: "Такого продукта не существует",
	},
	consts.CodeValidateError: {
		HttpCode: http.StatusBadRequest,
		Message: "Sorry, can't validate request",
		UserMessage: "Неверный формат параметров запроса",
	},
}


