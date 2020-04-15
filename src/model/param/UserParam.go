package param

// @Model         UpdateUserParam
// @Description   update user request parameter
// @Property      name string true "username"
type UpdateUserParam struct {
	Name string `json:"name" form:"name"`
}
