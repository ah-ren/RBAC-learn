package dto

// @Model         _UserDto
// @Description   user response
// @Property      id   integer true "user id"
// @Property      name string  true "username"
// @Property      role string  true "user role"
type UserDto struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}
