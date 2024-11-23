package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	auth "github.com/zjhM3l/go-e-commerce/backend/app/frontend/hertz_gen/frontend/auth"
	common "github.com/zjhM3l/go-e-commerce/backend/app/frontend/hertz_gen/frontend/common"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
