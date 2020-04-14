package exception

type ServerError struct {
	error
	Code    int32
	Message string
}

func NewError(code int32, message string) *ServerError {
	return &ServerError{
		Code:    code,
		Message: message,
	}
}
