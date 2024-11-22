namespace go frontend.order

include "api.thrift"
include "frontend/common.thrift"

// OrderService 提供订单相关接口服务
service OrderService {
    // 获取订单列表
    frontend.common.Empty OrderList(1: frontend.common.Empty request) (api.get="/order")
}