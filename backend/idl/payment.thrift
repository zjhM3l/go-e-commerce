namespace go payment

service PaymentService {
    // 执行支付
    ChargeResp Charge(1: ChargeReq request);
}

struct CreditCardInfo {
    1: string credit_card_number;        // 信用卡号
    2: i32 credit_card_cvv;              // CVV 码
    3: i32 credit_card_expiration_year;  // 信用卡到期年份
    4: i32 credit_card_expiration_month; // 信用卡到期月份
}

struct ChargeReq {
    1: double amount;                    // 支付金额
    2: CreditCardInfo credit_card;       // 信用卡信息
    3: string order_id;                  // 订单 ID
    4: i32 user_id;                      // 用户 ID
}

struct ChargeResp {
    1: string transaction_id;            // 交易 ID
}
