package errors

var (
	// Common errors
	OK                  = &HttpError{Code: 0, Message: "OK"}
	InternalServerError = &HttpError{Code: 10001, Message: "Internal server HttpError."}
	RouterNot           = &HttpError{Code: 10002, Message: "api route 404."}
)
