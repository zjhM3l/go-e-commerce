编码规范：gofmt + goimports

性能分析工具pprof
    CPU分析
    堆内存分析
    Groutine分析
    mutex锁分析
    block阻塞分析

开发框架
  1. Gorm：已经迭代了10年+的功能强大的ORM框架，在字节内部被广泛使用并且拥有非常丰富的开源扩展（OpenTelemetry监控扩展...）
  2. Kitex：字节内部的Golang微服务RPC框架，具有高性能、强可扩展的主要特点，支持多协议（GRPC、Thrift、Protobuf）并且拥有丰富的开源扩展（对接服务注册发现的注册中心，OpenTelemetry监控扩展...）
  3. Hertz：字节内部的 HTTP框架，参考了其他开源框架（fastHTTP）的优势，结合字节跳动内部的需求，具有高易用性、高性能、高扩展性 

性能提高策略：
  1. 去看高质量编程与性能调优实战里面
  2. GO三件套的Gorm性能提高两点策略，+ 分片库等