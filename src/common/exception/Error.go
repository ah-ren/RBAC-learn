package exception

import (
	"bytes"
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xcolor"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http/httputil"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

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
	Type    string   `json:"type"`
	Detail  string   `json:"detail"`
	Request []string `json:"request"`

	Filename string `json:"filename"`
	Function string `json:"function"`
	Line     int    `json:"line"`
	Content  string `json:"content"`
}

type Stack struct {
	Index    int
	Filename string
	Function string
	Pc       uintptr
	Line     int
	Content  string
}

func (s *Stack) String() string {
	return fmt.Sprintf("%s:%d (0x%x)\n\t%s: %s", s.Filename, s.Line, s.Pc, s.Function, s.Content)
}

func NewErrorDto(err interface{}, skip int, c *gin.Context, print bool) *ErrorDto {
	now := time.Now().Format(time.RFC3339)
	errType := fmt.Sprintf("%T", err)
	errDetail := fmt.Sprintf("%v", err)
	if e, ok := err.(error); ok {
		errDetail = e.Error()
	}

	// runtime
	stacks := stack(skip)
	filename := stacks[0].Filename
	function := stacks[0].Function
	line := stacks[0].Line
	content := stacks[0].Content
	if print {
		fmt.Println(xcolor.Yellow.Paint("\n[Panic Stack]"))
		for _, s := range stacks {
			fmt.Println(xcolor.Red.Paint(s.String()))
		}
		fmt.Println(xcolor.Yellow.Paint("[Panic Stack End]\n"))
	}

	// request
	requestBytes, _ := httputil.DumpRequest(c.Request, false)
	requestParams := strings.Split(string(requestBytes), "\r\n")
	request := make([]string, 0)
	for _, param := range requestParams {
		if strings.HasPrefix(param, "Authorization:") { // Authorization header
			request = append(request, "Authorization: *")
		} else if param != "" { // other param
			request = append(request, param)
		}
	}

	return &ErrorDto{
		Time:    now,
		Type:    errType,
		Detail:  errDetail,
		Request: request,

		Filename: filename,
		Function: function,
		Line:     line,
		Content:  content,
	}
}

func stack(skip int) []*Stack {
	out := make([]*Stack, 0)
	for i := skip; ; i++ {
		pc, filename, lineNumber, ok := runtime.Caller(i)
		if !ok {
			break
		}
		function := runtime.FuncForPC(pc).Name()
		_, function = filepath.Split(function)

		lineContent := "?"
		if filename != "" {
			if data, err := ioutil.ReadFile(filename); err == nil {
				lines := bytes.Split(data, []byte{'\n'})
				if lineNumber > 0 && lineNumber <= len(lines) {
					lineContent = string(bytes.TrimSpace(lines[lineNumber-1]))
				}
			}
		}
		out = append(out, &Stack{
			Index:    i,
			Filename: filename,
			Function: function,
			Pc:       pc,
			Line:     lineNumber,
			Content:  lineContent,
		})
	}

	return out
}
