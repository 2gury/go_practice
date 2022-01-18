package response

import "go_practice/8_clean_arch/internal/helpers/errors"

type Body map[string]interface{}

type Response struct {
	Code int `json:"code"`
	Error *errors.Error `json:"error,omitempty"`
	Body  *Body  `json:"body,omitempty"`
}
