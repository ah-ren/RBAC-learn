package param

// @Model         PolicyParam
// @Description   insert / delete policy request parameter
// @Property      role   string true "policy role"
// @Property      path   string true "policy path"
// @Property      method string true "policy method"
type PolicyParam struct {
	Role   string `json:"role"   form:"role"   binding:"required"`
	Path   string `json:"path"   form:"path"   binding:"required"`
	Method string `json:"method" form:"method" binding:"required"`
}

// @Model         RoleParam
// @Description   change user role request parameter
// @Property      role   string true "policy role"
type RoleParam struct {
	Role string `json:"role" form:"role" binding:"required"`
}
