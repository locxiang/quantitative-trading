package exceptions

var (
	// Common exceptions
	OK                  = &HttpError{Code: 0, Message: "OK"}
	InternalServerError = &HttpError{Code: 10001, Message: "Internal server HttpError."}

)
