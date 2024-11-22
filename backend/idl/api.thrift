namespace go api

// 扩展字段选项
struct FieldOptionsExtensions {
    1: optional string raw_body;            // 原始请求体
    2: optional string query;               // 查询参数
    3: optional string header;              // 请求头
    4: optional string cookie;              // Cookie
    5: optional string body;                // 请求体
    6: optional string path;                // 路径参数
    7: optional string vd;                  // 验证器
    8: optional string form;                // 表单数据
    9: optional string js_conv;             // JavaScript 转换
    10: optional string file_name;          // 文件名
    11: optional string none;               // 无
    31: optional string form_compatible;    // 表单兼容性
    32: optional string js_conv_compatible; // JavaScript 转换兼容性
    33: optional string file_name_compatible; // 文件名兼容性
    34: optional string none_compatible;    // 无兼容性
    101: optional string go_tag;            // Go 标签
}

// 扩展方法选项
struct MethodOptionsExtensions {
    1: optional string get;                 // GET 方法
    2: optional string post;                // POST 方法
    3: optional string put;                 // PUT 方法
    4: optional string delete;              // DELETE 方法
    5: optional string patch;               // PATCH 方法
    6: optional string options;             // OPTIONS 方法
    7: optional string head;                // HEAD 方法
    8: optional string any;                 // 任意方法
    31: optional string gen_path;           // 生成路径
    32: optional string api_version;        // API 版本
    33: optional string tag;                // 方法标签
    34: optional string name;               // 方法名称
    35: optional string api_level;          // 接口级别
    36: optional string serializer;         // 序列化方法
    37: optional string param;              // 公共参数
    38: optional string baseurl;            // 基础 URL
    39: optional string handler_path;       // 处理器路径
    61: optional string handler_path_compatible; // 处理器路径兼容性
}

// 扩展枚举值选项
struct EnumValueOptionsExtensions {
    1: optional i32 http_code;              // HTTP 状态码
}

// 扩展服务选项
struct ServiceOptionsExtensions {
    1: optional string base_domain;         // 基础域名
    31: optional string base_domain_compatible; // 基础域名兼容性
    32: optional string service_path;       // 服务路径
}

// 扩展消息选项
struct MessageOptionsExtensions {
    1: optional string reserve;             // 保留字段
}