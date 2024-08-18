## 目录结构
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

## 测试结果
```
{
  "invoice": {
    "id": "a6f18f9b-a0ca-402a-8225-48fd3e585f28",
    "orderId": "a0f9f7e2-dabe-41c5-9733-ac8367b65f47",
    "amount": 4999,
    "status": "paid"
  },
  "order": {
    "id": "a0f9f7e2-dabe-41c5-9733-ac8367b65f47",
    "customerId": "20ccaaec-fc7d-4d57-8d82-fbf9163d4017",
    "paymentId": "1a3f0aec-6518-4577-a3a0-18d4da353d08",
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
