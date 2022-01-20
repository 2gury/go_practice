package response

import "go_practice/9_clean_arch_db/internal/helpers/errors"

type Body map[string]interface{}

type Response struct {
	Code int `json:"code"`
	Error *errors.Error `json:"error,omitempty"`
	Body  *Body  `json:"body,omitempty"`
}
