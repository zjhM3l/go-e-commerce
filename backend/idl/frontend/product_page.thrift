namespace go frontend.product

include "api.thrift"
include "frontend/common.thrift"

// ProductReq 获取产品请求
struct ProductReq {
    1: i32 id (api.query="id") // 产品 ID
}

// SearchProductsReq 搜索产品请求
struct SearchProductsReq {
    1: string q (api.query="q") // 搜索关键词
}

// ProductService 提供产品相关接口服务
service ProductService {
    // 获取产品信息
    frontend.common.Empty GetProduct(1: ProductReq request) (api.get="/product")

    // 搜索产品
    frontend.common.Empty SearchProducs(1: SearchProductsReq request) (api.get="/search")
}