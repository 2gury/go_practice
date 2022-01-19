package request_reader

import (
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/schema"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
)

type QueryReader struct {
	decoder *schema.Decoder
}

func NewQueryReader() *QueryReader {
	return &QueryReader{
		decoder: schema.NewDecoder(),
	}
}

func (qr *QueryReader) Read(request interface{}, query map[string][]string) *errors.Error {
	qr.decoder.IgnoreUnknownKeys(true)
	err := qr.decoder.Decode(request, query)
	if err != nil {
		return errors.New(consts.CodeValidateError, err)
	}
	return nil
}

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
