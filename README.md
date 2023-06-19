
<div align=center>
    <h1>void-project</h1>
</div>
  

<div align=center>
    <a href="https://go.dev/doc/effective_go"><img src="https://img.shields.io/badge/GO-v1.20-blue"/></a>
    <a href="https://gin-gonic.com"><img src="https://img.shields.io/badge/Gin-v1.9.0-blue"/></a>
    <a href="https://gorm.io"><img src="https://img.shields.io/badge/GORM-v1.25.1-blue"/></a>
    <a href="https://redis.uptrace.dev"><img src="https://img.shields.io/badge/go--redis-v9.0.4-red"/></a>
    <a href="https://github.com/nhooyr.io/websocket"><img src="https://img.shields.io/badge/nhooyr.io/websocket-v1.8.7-green"/></a>
    <a href="https://github.com/golang-jwt/jwt"><img src="https://img.shields.io/badge/golang--jwt-v5-green"/></a>
</div>

## 介绍
void-project 是基于Gin + GORM + go-redis等构建的web应用集成后端架构，能够快速编写及实现web应用服务。符合Go语言定义的简单性开发原则设计哲学，同时兼顾了可拓展、易维护和规范化，在此中寻找平衡点。

分别参考了[golang-standards/project-layout](https://github.com/golang-standards/project-layout)社区约定俗成的标准布局，  
还参考传统mvc分层模式，  
以及Bob叔叔的[The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)架构。

## 特性
- 🚀得益于Go语言的性能优化，能够快速处理请求并响应。
- ⚙️无缝粘合Gin框架写法，自由定义包括：路由、中间件、响应体、模板引擎等等。
- ✅遵循RESTful API风格定义规范。
- ✅JWT Claims 鉴权/认证
- ⚙️多数据库连接(MySQL, SQLServer, SQLite...)，跨库查询等。
- ⚙️支持Redis的应用。
- ⚙️WebSocket接收/发送消息。
- 🔢遵循Go语言设计哲学：简单性，轻松编写前后分层的代码，可以将时间多放在业务逻辑的处理上，从而避免把时间浪费在底层基础和调整框架上。
- 🔢层次到模块间的代码规范化。

## TODO:
- 未来考虑实现自动化依赖注入容器。
- 接收请求和处理时的并发优化。
- 接下来准备引入OpenAPI(Swagger)接口文档解释库。由于目前api文件的生成库都是以注释方式生成的api定义文件，他这个注释加起来又麻烦占代码区域又大，滚动好几屏全是文档注释。所以我想等加入自动依赖注入后，按参数和返回方式生成，对于复杂结构体可以去容器里查找和分析。
- 一键数据库初始化/迁移。目前的数据库表生成比较麻烦，还需要到./cmd/install目录下手动编写映射结构体的迁移操作。
- 更多对Redis缓存的利用。
- 自己封装日志库，或引入Zap、logrus(第三方库有一点点的臃肿，我们需要的实际上有日志分级，信息格式化输出和写入文件就够了)。

## 目录结构
```
void-project
    ├── asset
    │    └── db
    ├── cmd
    │    ├── install
    │    ├── mark
    │    └── shiyan
    ├── config
    ├── initialize
    ├── internal
    │    ├── api
    │    │    ├── handler
    │    │    └── response
    │    ├── middleware
    │    ├── model
    │    ├── repository
    │    │    ├── driver
    │    │    ├── mysql
    │    │    ├── redis
    │    │    └── request
    │    ├── router
    │    ├── service
    │    └── view
    ├── pkg
    │    ├── jwt
    │    ├── logger
    │    └── md5
    ├── runtime
    │    └── log
    └── web
         ├── app
         ├── static
         │    ├── css
         │    └── js
         └── template

```

## 安装和使用
```
- 下载并安装Go
- Go版本 >= v1.18
- 安装MySQL,Redis等数据库/缓存
```
##### （1）无法访问Google系列网站
由于国内原因没法执行go get命令直接从google网站下载安装包管理库，推荐使用[goproxy.io](https://goproxy.io/zh/)或者[goproxy.cn](https://goproxy.cn/)设置免费的镜像代理。
操作方法：
```bash
# 第一步、将Go官方包管理工具(Go Modules)启用
go env -w GO111MODULE=on 
# 第二部、设置镜像代理的 GOPROXY 环境变量
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
```

##### （2）获取项目代码
下载项目代码和所需依赖库
```bash
# 克隆项目
git clone https://github.com/HyleaSoo/void-project.git
# 进入项目根目录
cd void-project

# 使用 go mod 安装go依赖包
go mod tidy
```

##### （3）运行示例项目
从根目录直接启动或是进入到cmd文件夹运行go run
```bash
# 方式1： 进入cmd文件夹，运行go run
cd cmd
go run .

# 方式2： 直接在根目录启动
go run cmd/main.go
```  

## 最后感谢您参与使用！  
<div>
    <span>🌌⚛️🔮🗡️✡️🏞️🎮</span>
    <label style="float:right;padding-right:30px;">—————— Hylea Soo<label>
</div>  




