namespace go frontend.home

include "api.thrift"
include "frontend/common.thrift"

// HomeService 提供主页相关接口服务
service HomeService {
    // 获取主页信息
    frontend.common.Empty Home(1: frontend.common.Empty request) (api.get="/")
}