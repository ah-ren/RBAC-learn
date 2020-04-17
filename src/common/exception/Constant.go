package exception

// Request
var (
	RequestParamError   = NewError(400, "request param error")
	ServerRecoveryError = NewError(500, "server unknown error")
)

// Authorization
var (
	UnAuthorizedError        = NewError(401, "unauthorized user")
	TokenExpiredError        = NewError(401, "token expired")
	InvalidRefreshTokenError = NewError(401, "invalid refresh token")
	CheckUserRoleError       = NewError(500, "failed to check user role")
	RoleHasNoPermissionError = NewError(403, "role has no permission")

	WrongPasswordError  = NewError(401, "wrong password")
	LoginError          = NewError(500, "login failed")
	RegisterError       = NewError(500, "register failed")
	RefreshTokenError   = NewError(500, "refresh token failed")
	LogoutError         = NewError(500, "logout failed")
	UpdatePasswordError = NewError(500, "update password failed")
)

// User
var (
	UserNotFoundError = NewError(404, "user not found")
	UserUpdateError   = NewError(500, "user update failed")
	UserDeleteError   = NewError(500, "user delete failed")
)

// Role
var (
	PolicyInsertError   = NewError(500, "insert policy failed")
	PolicyDeleteError   = NewError(500, "delete policy failed")
	PolicyNotFountError = NewError(404, "policy not found")
	PolicyExistedError  = NewError(409, "policy has existed")
	PolicySetRootError  = NewError(403, "set root role failed")
)
