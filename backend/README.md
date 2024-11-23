编码规范：gofmt + goimports

Makefile减少命令操作

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
  2. 接口定义-cwgo+thrift
    1. Interface Description Language：Thrift严格定义接口IDL
    2. 安装cwgo
    3. 基于IDL生成代码生成微服务应用-cwgo终端自动补全功能CloudWeGo All in one 代码生成工具
    添加自动补全组件终端cwgo然后可以查看
    创建server项目：
    mkdir -p app/service_name
    cwgo server --type RPC --module github.com/zjhM3l/go-e-commerce/app/service_name --service service_name --idl ../../idl/service.thrift
    <!-- 2. 使用Kitex生成带有脚手架的代码
    kitex -module github.com/zjhM3l/go-e-commerce -service checkoutservice idl/checkout.thrift
    3. 使用Hertz生成代码
    hz new -module github.com/zjhM3l/go-e-commerce -idl idl/auth.thrift
    如果有更新
    hz update -idl idl/user.thrift -->
  3. Kitex-服务注册与发现Etcd，Nacos，Zookeeper（最后用的ETCD）（见Go三件套笔记对应部分和代码示例）
  注册中心用consol，配置中心用etcd
    1. 服务端，服务启动相关代码位置，KitexInit中编写，复制官网consul相关示例，初始化consol，注册中心启动之后配置注册主键（注意要读取配置不要写死）
    2. 利用docker启动consol注册中心的容器
    3. 编写客户端代码，代码发现实例并且接口调用
    tips
    1. 建议使用Kitex的XDS扩展，多泳道做多套测试环境维护（多人开发）
    2. opentelemetry检测，打通gorm、kitex与hertz三件套，全流程链路追踪、检测、日志打通
  D:\studyANDworkFiles\bytedance24\kitex-registry-etcd-sample\registry-etcd\example\client\main.go
  4. 配置管理file(yaml, json, toml)/env/config，配置中心
    每个微服务下方conf，conf.go解析配置文件，dev/online/test分别对应开发生产测试三个环境配置
  5. GORM数据操作-数据存储：MySQL，Redis
  6. 微服务通信

第一版本时间紧，以后端为主，前端用hertz写的，没有做前后端分离，页面数据放在服务端生成
1. UI-Bootstrap，Fontawesome
2. 页面骨架：GolangTemplate生成 
  hertz提供了渲染数据的用法，结合golang的html/template和text/template
3. 操作
  1. frontend idl(为了练习，这里用的proto3)
  2. hertz快速生成服务端代码
  cd到app/frontend
  cwgo server --type HTTP --idl ../../idl/frontend/home.proto --service frontend -module github.com/zjhM3l/go-e-commerce/backend/app/frontend -I ../../idl
  业务逻辑biz
    dal数据库访问层
    handler=controller-只负责基本的参数校验和视图模型加载
    service负责主要的业务逻辑
  3. 改造首页接口，渲染html代码(hertz官方文档HTML渲染部分章节)
  4. 引入air热加载工具，方便调试
  5. bootstrap编写样式
  6. 布局优化


这里以home首页的逻辑再次加深理解一下整个框架流程
  1. 新增微服务（frontend）
  2. 微服务中新增功能（home）
    1. 注册路由到微服务frontend下的main.go
    2. hertz快速生成该功能服务端代码
  3. 业务实现
    1. 微服务下service文件夹中对应功能（home.go）中实现主体业务逻辑
    2. 微服务下handler问价家中对应功能（home_service.go）中为controller，只负责渲染页面和校验参数
    3. 无误后渲染页面

  用户服务
  1. 身份管理
    身份：身份证、邮箱、用户名密码...
    鉴权：服务端确认用户身份，但是用户（自然人）无法直接与服务端交互，委派给客户端（浏览器app）进行一定范围的数据资源操作
    授权结果：浏览器cookie和session，API凭证token
    Authorization授权：实现上述流程的过程
    Authentication鉴权：客户端被授予的身份鉴定和确认，对上述确认过的身份凭证进行解析和确认
    权限控制Permission control：对可执行的各种操作进行组合配置，成为权限列表，根据执行者的权限，操作在权限范围内允许执行，否则禁止
  2. Hertz相关中间件
    -hertz-contrib/sessions 管理session，基于cookie和redis两种实现方式，session最传统的会话技术，通过客户端传递sessionid来识别保存在服务端上的上下文信息
    -hertz-contrib/jwt 基于json的开放标准，基于访问令牌的身份认证技术，特别是单点登录的实现，实现了服务端的无状态，把状态信息保存在了客户端token中，需要注意对请求和数据的加密
    -hertz-contrib/paseto 安全无状态令牌的规范和参考实现，在鉴于jwt过于自由，容易出现漏洞和不完全算法的使用下，paseto提出了安全隐私易用性多语言跨平台的令牌方案
    -hertz-contrib/keyauth 工具库，帮助用户实现token的鉴权，协助在程序中对token的校验、统一处理、获取、在上下文中传递
  3. 登录注册页面
  4. 用户服务接口
  5. 鉴权和权限控制
    用户态写入session，根据session鉴权和权限控制
