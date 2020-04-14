package param

type TokenParam struct {
	RefreshToken string `json:"refresh-token" form:"refresh-token" binding:"required"`
	AccessToken  string `json:"access-token"  form:"access-token"  binding:"required"`
}
