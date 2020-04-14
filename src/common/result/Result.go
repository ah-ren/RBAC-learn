package result

import (
	"github.com/Aoi-hosizora/ahlib/xlinkedhashmap"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Result struct {
	Code    int32                         `json:"code"`
	Message string                        `json:"message"`
	Data    *xlinkedhashmap.LinkedHashMap `json:"data,omitempty"`
}

func Status(code int32) *Result {
	message := http.StatusText(int(code))
	if code == 200 {
		message = "success"
	} else if message == "" {
		message = "unknown"
	}
	return &Result{
		Code:    code,
		Message: strings.ToLower(message),
	}
}

func Ok() *Result {
	return Status(http.StatusOK)
}

func Error(code int32, message string) *Result {
	return Status(code).SetMessage(message)
}

func (r *Result) SetCode(code int32) *Result {
	r.Code = code
	return r
}

func (r *Result) SetMessage(message string) *Result {
	r.Message = strings.ToLower(message)
	return r
}

func (r *Result) SetData(data interface{}) *Result {
	r.Data = xlinkedhashmap.ObjectToLinkedHashMap(data)
	return r
}

func (r *Result) PutData(field string, data interface{}) *Result {
	if r.Data == nil {
		r.Data = xlinkedhashmap.NewLinkedHashMap()
	}
	r.Data.Set(field, data)
	return r
}

func (r *Result) JSON(c *gin.Context) {
	c.JSON(int(r.Code), r)
}

func (r *Result) XML(c *gin.Context) {
	c.XML(int(r.Code), r)
}