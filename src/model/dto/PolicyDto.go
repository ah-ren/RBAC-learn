package dto

// @Model         _PolicyDto
// @Description   policy response
// @Property      role   string true "policy role (sub)"
// @Property      path   string true "policy path (obj)"
// @Property      method string true "policy method (act)"
type PolicyDto struct {
	Role   string `json:"role"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
