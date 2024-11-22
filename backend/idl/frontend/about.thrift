namespace go frontend.about

include "api.thrift"
include "frontend/common.thrift"

// AboutService 提供 /about 接口的服务
service AboutService {
    // About 方法，接收一个 Empty 请求并返回一个 Empty 响应
    common.Empty About(1: common.Empty request) (api.post="/about")
}