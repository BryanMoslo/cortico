package models

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type APIError struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func NewResponse(statusCode int, message string) *Response {
	return &Response{
		Status:  statusCode,
		Message: message,
	}
}

func (r *Response) WithData(data interface{}) *Response {
	r.Data = data
	return r
}

func NewAPIError(status int, message string) *APIError {
	return &APIError{
		Status:  status,
		Message: message,
		Errors:  []string{},
	}
}
func (r *APIError) WithErrors(errors interface{}) *APIError {
	r.Errors = errors
	return r
}
