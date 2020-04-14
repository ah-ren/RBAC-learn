package dto

type UserDto struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}
