package param

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/config"
	"github.com/gin-gonic/gin"
	"strconv"
)

func BindId(c *gin.Context, field string) (uint32, bool) {
	uid, err := strconv.Atoi(c.Param(field))
	if err != nil {
		return 0, false
	}
	if uid <= 0 {
		return 0, false // <<<
	} else {
		return uint32(uid), true
	}
}

func BindPage(c *gin.Context, config *config.Config) (p int32, l int32) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page <= 0 {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "0"))
	if err != nil || limit <= 0 {
		limit = int(config.MetaConfig.DefPageSize)
	} else if limit > int(config.MetaConfig.MaxPageSize) {
		limit = int(config.MetaConfig.MaxPageSize)
	}
	return int32(page), int32(limit)
}
