package exception

import (
	"github.com/Aoi-hosizora/ahlib-gin-gorm/xgin"
)

func WrapValidationError(err error) *ServerError {
	isf := xgin.IsValidationFormatError(err)
	if isf {
		return RequestFormatError
	} else {
		return RequestParamError
	}
}
