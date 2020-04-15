package dto

// @Model         _TokenDto
// @Description   token response
// @Property      access-token  string true "access token"
// @Property      refresh-token string true "refresh token"
type TokenDto struct {
	AccessToken  string `json:"access-token"`
	RefreshToken string `json:"refresh-token"`
}
