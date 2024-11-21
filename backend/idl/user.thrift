namespace go user

service UserService {
    // 用户注册
    RegisterResp Register(1: RegisterReq request);
    // 用户登录
    LoginResp Login(1: LoginReq request);
}

struct RegisterReq {
    1: string email;               // 用户邮箱
    2: string password;            // 密码
    3: string confirm_password;    // 确认密码
}

struct RegisterResp {
    1: i32 user_id;                // 用户 ID
}

struct LoginReq {
    1: string email;               // 用户邮箱
    2: string password;            // 密码
}

struct LoginResp {
    1: i32 user_id;                // 用户 ID
}
