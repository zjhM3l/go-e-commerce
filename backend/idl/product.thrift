namespace go product

service ProductCatalogService {
    // 列出产品
    ListProductsResp ListProducts(1: ListProductsReq request);
    // 获取单个产品信息
    GetProductResp GetProduct(1: GetProductReq request);
    // 搜索产品
    SearchProductsResp SearchProducts(1: SearchProductsReq request);
}

struct ListProductsReq {
    1: i32 page;             // 页码
    2: i64 pageSize;         // 每页大小
    3: string categoryName;  // 分类名称
}

struct Product {
    1: i32 id;               // 产品 ID
    2: string name;          // 产品名称
    3: string description;   // 产品描述
    4: string picture;       // 产品图片链接
    5: double price;         // 产品价格
    6: list<string> categories; // 分类列表
}

struct ListProductsResp {
    1: list<Product> products; // 产品列表
}

struct GetProductReq {
    1: i32 id;               // 产品 ID
}

struct GetProductResp {
    1: Product product;      // 产品信息
}

struct SearchProductsReq {
    1: string query;         // 搜索关键词
}

struct SearchProductsResp {
    1: list<Product> results; // 搜索结果
}
