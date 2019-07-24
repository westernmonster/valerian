package http

import (
	"valerian/app/service/msm/model"
	"valerian/library/net/http/mars"
)

func scope(c *mars.Context) {
	var (
		err      error
		scopeMap map[int64]*model.Scope
		param    = new(struct {
			AppTreeID int64 `form:"app_tree_id" validate:"gt=0"`
		})
	)
	if err = c.Bind(param); err != nil {
		return
	}
	if scopeMap, err = svr.ServiceScopes(c, param.AppTreeID); err != nil {
		c.JSON(nil, err)
		return
	}
	data := make(map[string]interface{}, 2)
	data["service_tree_id"] = param.AppTreeID
	data["scopes"] = scopeMap
	c.JSON(data, nil)
}
