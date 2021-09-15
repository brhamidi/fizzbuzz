package http

// ResponseSuccess describes an generic API response for success
type ResponseSuccess struct {
	Data interface{} `json:"data"`
}

// ResponseError describes an generic API response for error
type ResponseError struct {
	Errs []Error `json:"errors,omitempty"`
} //@name ResponseError

// Error describes an error field in a Response
type Error struct {
	Message string `json:"message"`
} //@name Error

func NewResponseError(errs ...error) ResponseError {
	var resp ResponseError

	for _, err := range errs {
		resp.Errs = append(resp.Errs, Error{err.Error()})
	}
	return resp
}
