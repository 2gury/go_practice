package consts

type ErrorCode uint16

const (
	CodeBadRequest ErrorCode = 101 + iota
	CodeInternalError
	CodeValidateError
	CodeProductDoesNotExist
	CodeUserDoesNotExist
	CodeUserPasswordsDoNotMatch
	CodeWrongPasswords
	CodeStatusUnauthorized
	CodeUserNotConfirmation
)
