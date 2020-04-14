package param

type LoginParam struct {
	Id       uint32 `form:"id"       json:"id"       binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
