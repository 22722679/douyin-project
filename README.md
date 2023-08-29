# douyin-project
字节跳动青训营抖音项目

## 目录

### 项目架构
本项目采用单体架构，主要以Gin框架为主，采用MVC分层思想降低模块与模块之间的耦合度

### 项目技术栈

- GIN
- Viper
- Zap
- GORM
- MySQL
- 阿里云对象存储OOS

### 项目基础环境
- Go 1.20
- MySQL 5.7
- Gorm 2.x

### 文件目录
```go
├── conf  (配置文件信息)
│   └── config.go
├── controller (控制层)
│   ├── favorite.go
│   ├── Login.go
│   ├── publish.go
│   ├── register.go
│   └── user.go
├── msyql  (数据库操作层)
│   ├── feed.go
│   ├── mysql.go
│   └── user.go
├── middleware (中间件)
│   └── FTP (未开发)
├── model (模型层)
│   ├── favorite.go
│   ├── feed.go
│   ├── model.go
│   ├── publish.go
│   └── user.go
├── router (路由)
│   ├── router.go
│   └── videoresponse.go
├── service (业务逻辑层)
│   ├── favorite.go
│   ├── publish.go
│   └── user.go
├── oss
│   └── oss.go
└── main.go (主启动文件)
```

### 项目作者

### Git分支说明
- 主分支为完整项目
- 其他分支分别为部分功能模块（可能存在bug）

### 后续计划
- 升级为微服务项目，将功能模块抽取
