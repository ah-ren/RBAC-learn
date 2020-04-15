package controller

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/common/exception"
	"github.com/Aoi-hosizora/RBAC-learn/src/common/result"
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/Aoi-hosizora/RBAC-learn/src/database"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/dto"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/param"
	"github.com/Aoi-hosizora/RBAC-learn/src/service"
	"github.com/Aoi-hosizora/ahlib/xcondition"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/ahlib/xentity"
	"github.com/Aoi-hosizora/ahlib/xslice"
	"github.com/gin-gonic/gin"
)

type PolicyController struct {
	Config        *config.ServerConfig   `di:"~"`
	Mapper        *xentity.EntityMappers `di:"~"`
	UserService   *service.UserService   `di:"~"`
	CasbinService *service.CasbinService `di:"~"`
}

func NewPolicyController(dic *xdi.DiContainer) *PolicyController {
	ctrl := &PolicyController{}
	dic.MustInject(ctrl)
	return ctrl
}

// @Router              /v1/policy/role/{uid} [PUT]
// @Summary             Change user role
// @Security            Jwt
// @Tag                 Policy
// @Param               uid   path integer    true "user id"
// @Param               param body #RoleParam true "request parameter"
// @ResponseModel 200   #Result<Page<PolicyDto>>
func (r *PolicyController) SetRole(c *gin.Context) {
	uid, ok := param.BindId(c, "uid")
	roleParam := &param.RoleParam{}
	if err := c.ShouldBind(roleParam); err != nil || !ok {
		result.Error(exception.RequestParamError).JSON(c)
		return
	}

	user, ok := r.UserService.QueryById(uid)
	if !ok {
		result.Error(exception.UserNotFoundError).JSON(c)
		return
	}
	if user.Role == "root" {
		result.Error(exception.PolicySetRootError).JSON(c)
		return
	}
	user.Role = roleParam.Role

	status := r.UserService.Update(user)
	if status == database.DbNotFound {
		result.Error(exception.UserNotFoundError).JSON(c)
		return
	} else if status == database.DbFailed {
		result.Error(exception.UserUpdateError).JSON(c)
		return
	}

	result.Ok().JSON(c)
}

// @Router              /v1/policy [GET]
// @Summary             Query policy list
// @Template            Page
// @Security            Jwt
// @Tag                 Policy
// @ResponseModel 200   #Result<Page<PolicyDto>>
func (r *PolicyController) Query(c *gin.Context) {
	page, limit := param.BindPage(c, r.Config)
	total, policies := r.CasbinService.GetPolicies(limit, page)

	policiesDto := xcondition.First(r.Mapper.MapSlice(xslice.Sti(policies), &dto.PolicyDto{})).([]*dto.PolicyDto)
	result.Ok().SetPage(total, page, limit, policiesDto).JSON(c)
}

// @Router              /v1/policy [POST]
// @Summary             Insert policy list
// @Security            Jwt
// @Tag                 Policy
// @Param               param body #PolicyParam true "request parameter"
// @ResponseModel 200   #Result
func (r *PolicyController) Insert(c *gin.Context) {
	policyParam := &param.PolicyParam{}
	if err := c.ShouldBind(policyParam); err != nil {
		result.Error(exception.RequestParamError).JSON(c)
	}

	status := r.CasbinService.AddPolicy(policyParam.Role, policyParam.Path, policyParam.Method)
	if status == database.DbExisted {
		result.Error(exception.PolicyExistedError).JSON(c)
		return
	} else if status == database.DbFailed {
		result.Error(exception.PolicyInsertError).JSON(c)
		return
	}

	result.Ok().JSON(c)
}

// @Router              /v1/policy [DELETE]
// @Summary             Delete policy list
// @Security            Jwt
// @Tag                 Policy
// @Param               param body #PolicyParam true "request parameter"
// @ResponseModel 200   #Result
func (r *PolicyController) Delete(c *gin.Context) {
	policyParam := &param.PolicyParam{}
	if err := c.ShouldBind(policyParam); err != nil {
		result.Error(exception.RequestParamError).JSON(c)
	}

	status := r.CasbinService.DeletePolicy(policyParam.Role, policyParam.Path, policyParam.Method)
	if status == database.DbNotFound {
		result.Error(exception.PolicyNotFountError).JSON(c)
		return
	} else if status == database.DbFailed {
		result.Error(exception.PolicyDeleteError).JSON(c)
		return
	}

	result.Ok().JSON(c)
}
