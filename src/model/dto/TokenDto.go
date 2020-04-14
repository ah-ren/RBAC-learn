package dto

type TokenDto struct {
	RefreshToken string `json:"refresh-token"`
	AccessToken  string `json:"access-token"`
}
