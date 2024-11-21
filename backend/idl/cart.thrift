namespace go cart

service CartService {
    // 添加商品到购物车
    AddItemResp AddItem(1: AddItemReq request);
    // 获取购物车信息
    GetCartResp GetCart(1: GetCartReq request);
    // 清空购物车
    EmptyCartResp EmptyCart(1: EmptyCartReq request);
}

struct CartItem {
    1: i32 product_id;  // 商品 ID
    2: i32 quantity;    // 商品数量
}

struct AddItemReq {
    1: i32 user_id;     // 用户 ID
    2: CartItem item;   // 添加的商品
}

struct AddItemResp {
    // 空结构体，用于表示成功响应
}

struct EmptyCartReq {
    1: i32 user_id;     // 用户 ID
}

struct GetCartReq {
    1: i32 user_id;     // 用户 ID
}

struct GetCartResp {
    1: Cart cart;       // 返回的购物车信息
}

struct Cart {
    1: i32 user_id;                 // 用户 ID
    2: list<CartItem> items;        // 购物车中的商品列表
}

struct EmptyCartResp {
    // 空结构体，用于表示成功响应
}
