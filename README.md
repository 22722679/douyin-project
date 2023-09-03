# douyin-project
字节跳动青训营抖音项目

## 目录
- [项目技术栈](#项目技术栈)
- [项目基础环境](#项目基础环境)
- [安装步骤](#安装步骤)
- [文件目录](#文件目录)
- [项目作者](#项目作者)
- [使用到的技术](#使用到的技术)
- [版本控制](#版本控制)
- [鸣谢](#鸣谢)
### 项目技术栈

- gin
- Viper
- Zap
- GORM
- MySQL

### 项目基础环境
- Go 1.19
- MySQL 5.7
- Gorm 2.x
- sqlx v1.3.5

### 安装步骤
1. 下载源码
2. 配置相关服务器地址等相关参数
3. 启动服务
4. 在客户端配置相关地址服务端地址即可
```sh
git clone https://github.com/22722679/douyin-project.git
```
### 文件目录
```go
├── config(配置文件信息)
│   ├── config.go
│   └── info.go
├── controller(控制层)
│   ├── favorite.go
│   ├── Login.go
│   ├── publish.go
│   ├── register.go
│   └── user.go
├── middleware(中间件)
│   └── FTP (未开发)
├── model (模型层)
│   ├── favorite.go
│   ├── feed.go
│   ├── model.go
│   ├── publish.go
│   ├── registerInfo.go
│   ├── user.go
│   └── userInfo.go
├── msyql(数据库操作层)
│   ├── favorite.go
│   ├── feed.go
│   ├── mysql.go
│   ├── registerInfo.go
│   └── user.go
├── oss(云存储)
│   └── oss.go
├── router(路由)
│   ├── router.go
│   └── videoresponse.go
├── service(服务层)
│   ├── favorite.go
│   ├── publish.go
│   └── user.go
├── util(工具)
│   └── util.go
└── main.go (主启动文件)
```

### 项目作者
 张小虎
 
### 演示视频
[![Watch the video](https://lf3-static.bytednsdoc.com/obj/eden-cn/wthJoabvf_lm_tyvmahsWgpi/ljhwZthlaukjlkulzlp/images/introduce.png)](https://www.bilibili.com/video/BV1ow411S7vc/?vd_source=4cbd2d924b1efbce235b9f288896cf6f)

### 使用到的相关技术
框架相关：
- [Gin](https://gin-gonic.com/docs/)
- [Gorm](https://gorm.io/docs/)

数据库：
- [MySQL](https://dev.mysql.com/doc/)
### 版本控制

该项目使用Git进行版本管理。您可以在repository参看当前可用版本。

### 鸣谢

- [字节跳动后端青训营](https://youthcamp.bytedance.com/)
