namespace go frontend.checkout

include "api.thrift"
include "frontend/common.thrift"

// CheckoutReq 结算请求
struct CheckoutReq {
    1: string email (api.form="email")                       // 用户邮箱
    2: string firstname (api.form="firstname")               // 用户名
    3: string lastname (api.form="lastname")                 // 用户姓
    4: string street (api.form="street")                     // 街道地址
    5: string zipcode (api.form="zipcode")                   // 邮政编码
    6: string province (api.form="province")                 // 省/州
    7: string country (api.form="country")                   // 国家
    8: string city (api.form="city")                         // 城市
    9: string card_num (api.form="cardNum")                  // 卡号
    10: i32 expiration_month (api.form="expirationMonth")    // 到期月份
    11: i32 expiration_year (api.form="expirationYear")      // 到期年份
    12: i32 cvv (api.form="cvv")                             // CVV
    13: string payment (api.form="payment")                  // 支付方式
}

// CheckoutService 提供结算相关接口服务
service CheckoutService {
    // 结算
    common.Empty Checkout(1: CheckoutReq request) (api.get="/checkout")

    // 等待结算
    common.Empty CheckoutWaiting(1: common.Empty request) (api.post="/checkout/waiting")

    // 获取结算结果
    common.Empty CheckoutResult(1: common.Empty request) (api.get="/checkout/result")
}