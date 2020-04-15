package param

type RegisterParam struct {
	Password string `json:"password" form:"password" binding:"required"`
	Name     string `json:"name"     form:"name"`
}
