编码规范：gofmt + goimports

性能分析工具pprof
    CPU分析
    堆内存分析
    Groutine分析
    mutex锁分析
    block阻塞分析

开发框架https://www.bilibili.com/video/BV1bf421o7NM?spm_id_from=333.788.videopod.sections&vd_source=c526cc536bc0044743ba2f78dee5f33e里面3:35出现过
  1. Gorm：已经迭代了10年+的功能强大的ORM框架，在字节内部被广泛使用并且拥有非常丰富的开源扩展（OpenTelemetry监控扩展...）
  2. Kitex：字节内部的Golang微服务RPC框架，具有高性能、强可扩展的主要特点，支持多协议（GRPC、Thrift、Protobuf）并且拥有丰富的开源扩展（对接服务注册发现的注册中心，OpenTelemetry监控扩展...）
  3. Hertz：字节内部的 HTTP框架，参考了其他开源框架（fastHTTP）的优势，结合字节跳动内部的需求，具有高易用性、高性能、高扩展性 

  参考来源：https://juejin.cn/post/7268539624560853051
  当 Hertz 和 Kitex 结合使用时，Hertz 充当了 HTTP 服务器的角色，而 Kitex 充当了 RPC 框架的角色。Hertz 提供了 HTTP 服务的能力，可以接收来自客户端的 HTTP 请求，并将请求转发给 Kitex 处理。Kitex 则提供了 RPC 调用的能力，可以处理来自 Hertz 的请求，并将结果返回给 Hertz。这种结合使用的方式可以让开发者更加方便地构建高性能、易扩展、低时延的微服务。
  为了更好地理解 Hertz 和 Kitex 的结合使用，我们可以想象一个场景：假设我们正在开发一个在线商城应用程序，需要提供商品查询、下单、支付等功能。我们可以使用 Hertz 提供 HTTP 服务，接收来自客户端的 HTTP 请求，并将请求转发给 Kitex 处理。Kitex 可以处理来自 Hertz 的请求，并调用后端服务进行商品查询、下单、支付等操作。最后，Kitex 将结果返回给 Hertz，Hertz 再将结果返回给客户端。这样，我们就可以通过 Hertz 和 Kitex 结合使用，构建一个高性能、易扩展、低时延的在线商城应用程序。

性能提高策略：
  1. 去看高质量编程与性能调优实战里面
  2. GO三件套的Gorm性能提高两点策略，+ 分片库等

流程工具
  1. API测试工具Postman API
  2. 接口定义
    1. Interface Description Language：Thrift严格定义接口IDL
    2. 使用Kitex生成带有脚手架的代码
    kitex -module github.com/zjhM3l/go-e-commerce -service checkoutservice idl/checkout.thrift
    3. 使用Hertz生成代码
    hz new -module github.com/zjhM3l/go-e-commerce -idl idl/auth.thrift
    如果有更新
    hz update -idl idl/user.thrift
  3. Kitex-服务注册与发现Etcd，Nacos，Zookeeper（最后用的ETCD）（见Go三件套笔记对应部分和代码示例）
  注册中心用consol，配置中心用etcd
  D:\studyANDworkFiles\bytedance24\kitex-registry-etcd-sample\registry-etcd\example\client\main.go
  4. 建议使用Kitex的XDS扩展，多泳道做多套测试环境维护（多人开发）
  5. opentelemetry检测，打通gorm、kitex与hertz三件套，全流程链路追踪、检测、日志打通
  6. 数据存储：MySQL，Redis
