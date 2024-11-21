namespace go checkout

include "payment.thrift"

service CheckoutService {
    // 执行结算
    CheckoutResp Checkout(1: CheckoutReq request);
}

struct Address {
    1: string street_address;  // 街道地址
    2: string city;            // 城市
    3: string state;           // 州/省
    4: string country;         // 国家
    5: string zip_code;        // 邮政编码
}

struct CheckoutReq {
    1: i32 user_id;                 // 用户 ID
    2: string firstname;            // 名
    3: string lastname;             // 姓
    4: string email;                // 用户邮箱
    5: Address address;             // 收货地址
    6: payment.CreditCardInfo credit_card; // 信用卡信息
}

struct CheckoutResp {
    1: string order_id;             // 订单 ID
    2: string transaction_id;       // 交易 ID
}
