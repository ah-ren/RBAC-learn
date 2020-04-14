package exception

// Request
var (
	RequestParamError  = NewError(400, "request param error")
	RequestFormatError = NewError(400, "request format error")
)

// Authorization
var (
	UnAuthorizedError        = NewError(401, "unauthorized user")
	TokenExpiredError        = NewError(401, "token expired")
	InvalidRefreshTokenError = NewError(401, "invalid refresh token")

	WrongPasswordError = NewError(401, "wrong password")
	LoginError         = NewError(500, "login failed")
	RefreshTokenError  = NewError(500, "refresh token failed")
)
