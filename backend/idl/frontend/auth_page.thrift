namespace go frontend.auth

include "api.thrift"
include "frontend/common.thrift"

// 注册请求消息
struct RegisterReq {
    1: string email (api.form="email")               // 用户邮箱
    2: string password (api.form="password")         // 用户密码
    3: string confirm_password (api.form="confirm_password") // 确认密码
}

// 登录请求消息
struct LoginReq {
    1: string email (api.form="email")               // 用户邮箱
    2: string password (api.form="password")         // 用户密码
    3: string next (api.query="next")                // 下一步跳转链接
}

// AuthService 提供认证相关接口服务
service AuthService {
    // 用户注册接口
    common.Empty register(1: RegisterReq request) (api.post="/auth/register")

    // 用户登录接口
    common.Empty login(1: LoginReq request) (api.post="/auth/login")

    // 用户登出接口
    common.Empty logout(1: common.Empty request) (api.post="/auth/logout")
}