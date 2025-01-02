# Gin-ProHub项目
## 介绍
本项目为本人心血来潮想使用一个后端框架来开发个人主页+博客，但苦于Java只使用基础语法 不学习Spring等框架无法直接编写，于是使用Go来完成后端内容
## 结构
backend/
├── cmd/
│   └── gin-prohub.go      # 主程序入口
├── config/
│   ├── config.go          # 配置加载
│   ├── config.yaml        # 配置文件
│   └── log.yaml          # 日志配置
├── middleware/           # 新建中间件目录
│   ├── auth.go          # 认证中间件
│   └── cors.go          # 跨域中间件
├── models/              # 数据模型
│   ├── Posts.go
│   └── User.go
├── routes/             # 路由
│   ├── auth_routes.go
│   ├── routes.go
│   └── User_routes.go
├── services/          # 业务逻辑
│   └── user_service.go
└── utils/            # 工具包
    ├── errors/       # 错误处理
    │   └── errors.go
    ├── logger/       # 日志处理
    │   └── logger.go
    └── response/     # 响应封装
        └── response.go