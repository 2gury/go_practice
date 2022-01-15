package response

type Body map[string]interface{}

type Response struct {
	Code int     `json:"code"`
	Error string `json:"error,omitempty"`
	Body *Body   `json:"body,omitempty"`
}
