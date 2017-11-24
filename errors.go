package errors

var (
	InvalidParameter = Error{400, "InvalidParameter", "A parameter specified in a request is not valid, is unsupported, or cannot be used."}
	AccessDenied     = Error{403, "AccessDenied", "Access Denied"}
	InternalError    = Error{500, "InternalError", "An internal error has occurred. Retry your request, but if the problem persists, contact us with details by posting a message on the service forums."}
)
