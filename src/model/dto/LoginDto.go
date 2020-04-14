package dto

type LoginDto struct {
	User         *UserDto `json:"user"`
	AccessToken  string   `json:"access-token"`
	RefreshToken string   `json:"refresh-token"`
}
