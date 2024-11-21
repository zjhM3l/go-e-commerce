namespace go auth

service AuthService {
    // 用于发放令牌
    DeliverTokenResp DeliverTokenByRPC(1: DeliverTokenReq request);
    // 用于验证令牌
    VerifyResp VerifyTokenByRPC(1: VerifyTokenReq request);
}

struct DeliverTokenReq {
    1: i32 user_id; // 用户 ID
}

struct VerifyTokenReq {
    1: string token; // 待验证的令牌
}

struct DeliverTokenResp {
    1: string token; // 返回的令牌
}

struct VerifyResp {
    1: bool res; // 验证结果，true 为成功
}
