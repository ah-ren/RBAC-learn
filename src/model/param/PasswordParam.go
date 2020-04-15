package param

type PasswordParam struct {
	Password string `json:"password" form:"password" binding:"required"`
}
