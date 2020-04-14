package param

type LoginParam struct {
	Id       uint32 `json:"id"       form:"id"       binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
