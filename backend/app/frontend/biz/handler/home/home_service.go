package home

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/zjhM3l/go-e-commerce/backend/app/frontend/biz/service"
	"github.com/zjhM3l/go-e-commerce/backend/app/frontend/biz/utils"
	home "github.com/zjhM3l/go-e-commerce/backend/app/frontend/hertz_gen/frontend/home"
)

// Home .
// @router / [GET]
func Home(ctx context.Context, c *app.RequestContext) {
	var err error
	var req home.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewHomeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	// controller只负责渲染页面和校验参数，具体的业务逻辑在service中实现
	c.HTML(consts.StatusOK, "home.tmpl", resp)
}
