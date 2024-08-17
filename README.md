```
.
├── Makefile  依赖按照（grpc、openapi）
├── baskets   购物车服务
│   ├── basketspb （proto 定义）
│   │   ├── api.pb.go
│   │   ├── api.pb.gw.go
│   │   ├── api.proto
│   │   ├── api_grpc.pb.go
│   │   ├── messages.pb.go
│   │   └── messages.proto
│   ├── buf.gen.yaml （自动生成 proto 相关代码的 yml 文件）
│   ├── buf.yaml
│   ├── generate.go   
│   ├── internal         
│   │   ├── application（应用层实现 CRUD）
│   │   ├── domain（领域模型）
│   │   ├── grpc （grpc 服务间通信）
│   │   ├── logging（日志格式化）
│   │   ├── postgres
│   │   └── rest（RESTApi）
│   └── module.go（模块化）
├── cmd（服务启动入口）
│   └── mallbots
│       ├── main.go（start app with module）
│       └── monolith.go（app start config）
customers(用户服务)
│   ├── buf.gen.yaml
│   ├── buf.yaml
│   ├── customerspb
│   ├── generate.go
│   ├── internal
│   └── module.go
├── depot（仓库服务）核心
│   ├── buf.gen.yaml
│   ├── buf.yaml
│   ├── depotpb
│   ├── generate.go
│   ├── internal
│   └── module.go
├── docker（docker 启动+database 数据库创建）
│   ├── Dockerfile
│   ├── database
│   └── wait-for
├── docker-compose.yml
├── docs
│   ├── ADL
│   ├── Diagrams（项目流程图）
│   └── EventStorming
├── go.mod
├── go.sum
├── internal
│   ├── config
│   ├── ddd
│   ├── logger
│   ├── monolith
│   ├── rpc
│   ├── tools.go
│   ├── waiter
│   └── web
├── notifications（通知服务）
│   ├── buf.gen.yaml
│   ├── buf.yaml
│   ├── generate.go
│   ├── internal
│   ├── module.go
│   └── notificationspb
├── ordering（订单服务）
│   ├── buf.gen.yaml
│   ├── buf.yaml
│   ├── generate.go
│   ├── internal
│   ├── module.go
│   └── orderingpb
├── payments（支付服务）
│   ├── buf.gen.yaml
│   ├── buf.yaml
│   ├── generate.go
│   ├── internal
│   ├── module.go
│   └── paymentspb
└── stores（商店管理服务）
    ├── buf.gen.yaml
    ├── buf.yaml
    ├── features
    ├── generate.go
    ├── internal
    ├── module.go
    ├── stores_test.go
    └── storespb
```


服务开发流程
1. 定义 proto 文件
2. 定义 buf.yaml 和 buf.gen.yaml 文件
3. 定义 api.annotations.yaml、api.openapi.yaml、index.html 文件
4. 使用 buf generate 生成 api.swagger.json、api_grpc.pb.go、api.pb.go、api.pb.gw.go、message.pb.go 文件
5. 编写 gateway 反向代理器、编写 swagger.go 暴露 baskets-spec 接口
6. 开发 internal 
   1. application
   2. domain
   3. grpc
   4. logging
   5. postgres
   6. models（可选）

```
{
  "order": {
    "id": "a2065089-836a-4db6-8bb0-860c6b7e1ac5",
    "customerId": "fff2d41a-a485-4008-b578-747732ae1089",
    "paymentId": "90b2e735-13d1-4ae1-852a-99ead33492b0",
    "items": [
      {
        "storeId": "bd9e26cd-d861-4b74-862f-06e0825c7af3",
        "productId": "557e2f05-b10a-475c-baff-2cea8a86d8df",
        "storeName": "Waldorf Books",
        "productName": "EDA with Golang",
        "price": 49.99,
        "quantity": 100
      }
    ],
    "status": "completed"
  }
}
```
