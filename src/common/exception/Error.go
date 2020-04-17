package exception

type Error struct {
	error
	Code    int32
	Message string
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
