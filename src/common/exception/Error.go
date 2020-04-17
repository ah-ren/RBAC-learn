package exception

type Error struct {
	Code    int32
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code int32, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

type ErrorDto struct {
	Time    string   `json:"time"`
	File    string   `json:"file"`
	Line    int      `json:"line"`
	Func    string   `json:"func"`
	Detail  string   `json:"detail"`
	Request []string `json:"request"`
}
