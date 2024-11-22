namespace go frontend.cart

include "api.thrift"
include "frontend/common.thrift"

// AddCartReq 添加购物车请求
struct AddCartReq {
    1: i32 product_id (api.form="productId")   // 产品 ID
    2: i32 product_num (api.form="productNum") // 产品数量
}

// CartService 提供购物车相关接口服务
service CartService {
    // 添加商品到购物车
    common.Empty AddCartItem(1: AddCartReq request) (api.post="/cart")

    // 获取购物车信息
    common.Empty GetCart(1: common.Empty request) (api.get="/cart")
}