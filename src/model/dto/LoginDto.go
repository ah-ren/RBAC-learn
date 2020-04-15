package dto

// @Model         _LoginDto
// @Description   login response
// @Property      user          object(#_UserDto) true "user response"
// @Property      access-token  string            true "access token"
// @Property      refresh-token string            true "refresh token"
type LoginDto struct {
	User         *UserDto `json:"user"`
	AccessToken  string   `json:"access-token"`
	RefreshToken string   `json:"refresh-token"`
}
