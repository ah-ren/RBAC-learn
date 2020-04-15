package param


// @Model         LoginParam
// @Description   login request parameter
// @Property      id       integer true "user id"
// @Property      password string  true "password"
type LoginParam struct {
	Id       uint32 `json:"id"       form:"id"       binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// @Model         RegisterParam
// @Description   register request parameter
// @Property      password string true  "password"
// @Property      name     string false "username"
type RegisterParam struct {
	Password string `json:"password" form:"password" binding:"required"`
	Name     string `json:"name"     form:"name"`
}

// @Model         TokenParam
// @Description   refresh token request parameter
// @Property      refresh-token string true "refresh token"
// @Property      access-token  string true "access token"
type TokenParam struct {
	RefreshToken string `json:"refresh-token" form:"refresh-token" binding:"required"`
	AccessToken  string `json:"access-token"  form:"access-token"  binding:"required"`
}

// @Model         PasswordParam
// @Description   reset password request parameter
// @Property      password string true "new password"
type PasswordParam struct {
	Password string `json:"password" form:"password" binding:"required"`
}
