package errors

type ErrorResponse struct {
	statusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Resource   string `json:"resource,omitempty"`
	RequestID  string `json:"request_id"`
	RawError   string `json:"error"`
}

func NewErrorResponse(requestID, resource string, err Error, errs ...error) *ErrorResponse {
	er := &ErrorResponse{
		statusCode: err.Code,
		Code:       err.Name,
		Message:    err.Message,
		Resource:   resource,
		RequestID:  requestID,
	}
	if len(errs) > 0 {
		er.RawError = errs[0].Error()
	}

	return er
}

// implements gogo.StatusCoder interface
func (er *ErrorResponse) StatusCode() int {
	return er.statusCode
}

func (er *ErrorResponse) Error() string {
	if er.RawError != "" {
		return "[" + er.Code + "] " + er.RawError
	}

	return "[" + er.Code + "] " + er.Message
}
