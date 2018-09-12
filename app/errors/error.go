package errors



type HttpError struct {
	Code    int
	Message string
}

func (e HttpError) Error() string {
	return e.Message
}
