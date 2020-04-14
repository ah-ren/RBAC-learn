package dto

type LoginDto struct {
	User  *UserDto `json:"user"`
	Token string   `json:"token"`
}
