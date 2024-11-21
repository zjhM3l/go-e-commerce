namespace go order

include "cart.thrift"

service OrderService {
    // 下订单
    PlaceOrderResp PlaceOrder(1: PlaceOrderReq request);
    // 列出用户订单
    ListOrderResp ListOrder(1: ListOrderReq request);
    // 标记订单已支付
    MarkOrderPaidResp MarkOrderPaid(1: MarkOrderPaidReq request);
}

struct Address {
    1: string street_address;  // 街道地址
    2: string city;            // 城市
    3: string state;           // 州/省
    4: string country;         // 国家
    5: i32 zip_code;           // 邮政编码
}

struct PlaceOrderReq {
    1: i32 user_id;                    // 用户 ID
    2: string user_currency;           // 用户货币类型
    3: Address address;                // 收货地址
    4: string email;                   // 用户邮箱
    5: list<OrderItem> order_items;    // 订单商品列表
}

struct OrderItem {
    1: cart.CartItem item;             // 购物车中的商品
    2: double cost;                    // 商品价格
}

struct OrderResult {
    1: string order_id;                // 订单 ID
}

struct PlaceOrderResp {
    1: OrderResult order;              // 下单结果
}

struct ListOrderReq {
    1: i32 user_id;                    // 用户 ID
}

struct Order {
    1: list<OrderItem> order_items;    // 订单商品列表
    2: string order_id;                // 订单 ID
    3: i32 user_id;                    // 用户 ID
    4: string user_currency;           // 用户货币类型
    5: Address address;                // 收货地址
    6: string email;                   // 用户邮箱
    7: i32 created_at;                 // 创建时间
}

struct ListOrderResp {
    1: list<Order> orders;             // 用户订单列表
}

struct MarkOrderPaidReq {
    1: i32 user_id;                    // 用户 ID
    2: string order_id;                // 订单 ID
}

struct MarkOrderPaidResp {
    // 空结构体，用于表示成功响应
}
