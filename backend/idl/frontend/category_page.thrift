namespace go frontend.category

include "api.thrift"
include "frontend/common.thrift"

// CategoryReq 类别请求
struct CategoryReq {
    1: string category (api.path="category") // 类别名称
}

// CategoryService 提供类别相关接口服务
service CategoryService {
    // 获取类别信息
    common.Empty Category(1: CategoryReq request) (api.get="/category/:category")
}