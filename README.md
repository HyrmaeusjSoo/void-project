
<div align=center>
    <h1>void-project</h1>
</div>
  

<div align=center>
    <img src="https://img.shields.io/badge/version-1.0.28-05e5a5">
    <a href="https://go.dev/doc/effective_go"><img src="https://img.shields.io/badge/Go-v1.21-blue"/></a>
    <a href="https://gin-gonic.com"><img src="https://img.shields.io/badge/Gin-v1.10.0-blue"/></a>
    <a href="https://gorm.io"><img src="https://img.shields.io/badge/GORM-v1.25.12-blue"/></a>
    <a href="https://redis.uptrace.dev"><img src="https://img.shields.io/badge/go--redis-v9.6.1-red"/></a>
    <a href="https://github.com/nhooyr/websocket"><img src="https://img.shields.io/badge/nhooyr.io/websocket-v1.8.11-green"/></a>
    <a href="https://github.com/golang-jwt/jwt"><img src="https://img.shields.io/badge/golang--jwt-v5.2.1-green"/></a>
</div>

<br/>

## 介绍
void-project 是基于Gin + GORM + go-redis等构建的Web应用集成后端架构，能够快速编写及实现Web应用服务。符合Go语言定义的简单性原则设计哲学进行开发，同时兼顾了可拓展、易维护和规范化，在此间寻找平衡点。

分别参考了[golang-standards/project-layout](https://github.com/golang-standards/project-layout)社区约定俗成的标准布局，  
还参考传统mvc分层模式，  
以及Bob叔叔的[The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)架构。

<br/>

## Vue前端请求示例
后端的示例已经尽量把Go和几个框架的特性都用起来。  
与之配套的前端Vue简单的请求示例地址：[void-nebula](https://github.com/HyrmaeusjSoo/void-nebula)

<br/>

## 特性
- 🚀得益于Go语言的性能优化，能够快速处理请求并响应。
- 🚀无缝粘合Gin框架写法，自由定义包括：路由、中间件、响应体、模板引擎等等，快速实现Web应用。
- ✅遵循RESTful API风格定义规范。
- ✅标准JWT Claims 鉴权/认证。
- ⚙️多数据库连接(MySQL, SQLServer, SQLite...)，跨库查询、分页、自定义Scope等快捷操作。
- ⚙️支持Redis的应用。
- ⚙️WebSocket接收/发送消息。
- 📃使用简单的自定义日志。实现控制台输出和写入文件，简单的日志分级，日志文件分层分日期方式记录。当然，同时也可以引入第三方日志库。
- 🔢遵循Go语言设计哲学：简单性原则，轻松编写前后分层的代码，可以将时间多放在业务逻辑的处理上，从而避免把时间浪费在底层基础和调整框架上。
- 🔢层次到模块间的代码规范化。

<br/>

## TODO:
- 未来考虑实现自动化依赖注入容器。
- 接收请求和处理时的并发优化。
- 接下来准备引入OpenAPI(Swagger)接口文档解释库。由于目前api文件的生成库都是以注释方式生成的api定义文件，他这个注释加起来又麻烦占代码区域又大，滚动好几屏全是文档注释。所以我想等加入自动依赖注入后，按参数和返回方式生成，对于复杂结构体可以去容器里查找和分析。
- 一键数据库初始化/迁移。目前的数据库表生成比较麻烦，还需要到./cmd/install目录下手动编写映射结构体的迁移操作。
- 考虑将自定义logger结合到Go1.21.0内置的slog包来封装日志，写入json格式，以便后续读取和分析。

<br/>

## 目录结构
```
──────────────────begin──────────────────
void-project
    ├── asset
    │    ├── database
    │    └── json
    ├── cmd
    │    ├── client
    │    ├── install
    │    ├── mark
    │    └── server
    ├── config
    ├── global
    ├── initialize
    ├── internal
    │    ├── api
    │    │    ├── handler
    │    │    ├── request
    │    │    └── response
    │    │         └── apierr
    │    ├── middleware
    │    ├── model
    │    │    └── base
    │    ├── repository
    │    │    ├── driver
    │    │    ├── mysql
    │    │    │    └── scope
    │    │    ├── redis
    │    │    ├── request
    │    │    └── sqlite
    │    ├── router
    │    ├── service
    │    └── view
    ├── pkg
    │    ├── bcrypt
    │    ├── clicolor
    │    ├── jwt
    │    ├── logger
    │    │    └── slog
    │    ├── md5
    │    ├── necromancy
    │    ├── translation
    │    └── types
    │         ├── composite
    │         └── primitive
    ├── runtime
    │    ├── log
    │    │    ├── error
    │    │    ├── server
    │    │    └── sql
    │    └── slog
    └── web
         ├── app
         ├── static
         │    ├── css
         │    ├── image
         │    │    └── icon
         │    └── js
         ├── template
         └── upload
              └── 模块/年/月/日/
───────────────────end───────────────────
```

<br/>

## 获取和使用
*从1.21.0开始，go.mod中的版本号将作为编译时的最小Go语言版本，这是官方新的约束规则。
```
- 下载并安装Go
- Go版本 >= v1.21.0 *项目中使用了Go的新版本新特性
- 安装MySQL,Redis等数据库/缓存
```
##### （1）无法访问Google系列网站
由于国内原因没法执行go get命令直接从google网站下载安装包管理库，推荐使用[goproxy.io](https://goproxy.io/zh/)或者[goproxy.cn](https://goproxy.cn/)设置免费的镜像代理。
操作方法：
```Shell
# 第一步、将Go官方包管理工具(Go Modules)启用
go env -w GO111MODULE=on 
# 第二步、设置镜像代理的 GOPROXY 环境变量
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
```

##### （2）获取项目代码
下载项目代码和所需依赖库
```Shell
# 克隆项目
git clone https://github.com/HyrmaeusjSoo/void-project.git
# 进入项目根目录
cd void-project

# 使用 go mod 安装go依赖包
go mod tidy
```

##### （3）运行示例项目
从根目录直接启动或是进入到cmd文件夹运行go run
```Shell
# 方式1： 进入cmd下的server、install等目录，运行go run
cd cmd/server
go run .

# 方式2： 直接在根目录启动
go run cmd/server/main.go

# OK!
```

<br/>

## 编译
go编译有很多参数和方式，最好进入到对应的目录内根据需要添加参数进行编译。下面为编译示例：
##### 正常编译
正常编译时，不做其他特殊处理，进入入口目录直接执行go build命令。
```Shell
# 1. 进入cmd目录内server、install等目录
cd cmd/server

# 2. 执行编译命令
go build -ldflags "-s -w" -v -trimpath
```
##### 交叉编译
交叉编译首先要使用go命令设置语言环境中对应的目标系统和cpu架构位数等。  
 ！因为在windows系统cmd和powershell设置语言环境的命令不一样，所以这里统一用go env的设置方式，之后可以随意改回来。  
 简短方便的编译命令自行查找吧
```Shell
# 1. 设置编译目标信息
go env -w GOOS=linux    # 设置目标系统
go env -w GOARCH=amd64  # 设置目标cpu架构及位数
go env -w CGO_ENABLED=0 # 关闭cgo，某些系统下的cgo都不一样。目前没用到cgo

# 2. 接下来在cmd目录下的某项内执行编译命令，就自动打包成目标系统的可执行文件了
go build -ldflags "-s -w" -v -trimpath
```

##### LSP-DAP
在Terminal工作流写代码时需要配置DAP，Go默认使用的是%GOPATH%/bin/dlv。  
vim、helix等编辑器选择binary方式启动debugger需要手动编译
```Shell
# -N: 禁止编译器优化
# -l: 禁止内联函数
# -v: 打印包编译顺序(可选)
go build -gcflags "-N -l" -v
```

<br/>

#### 特别提示
！！！2023-07-15更新之后不用改写代码了，只要在config/system.json或者.yaml配置文件内加入mode节点为release就是生产环境！！！  

生产环境中应该简单直观最好，因此建议在生产环境设置Gin模式为‘发布模式’，并且在Gin和GORM的日志配置选项中禁用色彩打印。
```GO
// 1.设置Gin启动模式为发布模式
//    在cmd/main.go文件里，main方法内。Server的模式设置为ReleaseMode。
gin.SetMode(gin.ReleaseMode)

// 2.禁用Gin的彩色日志
//    在initialize/initialize.go文件里，InitServerLog方法内。
//    Gin日志中的‘强制控制台色彩’删掉或注释。
gin.ForceConsoleColor() // 删掉或注释或改为 gin.DisableConsoleColor()

// 3.禁用GORM日志的彩色打印
//    internal/repository/driver/db_conn.go文件里。
//    所有初始化连接方法内logger.Config中设置项为Colorful:true全部改为false
Logger: logger.New(
    log.NewSQLLogger(),
    logger.Config{
        ...
        Colorful: true,  // 将所有此设置项改为false
    },
),

```  

<br/>

## 测试环境与部署
测试环境直接执行文件启动更方便，但生产环境推荐nginx作为前置/反向代理服务器等等等更安全/保守的方式。
##### 主要目录
```Shell
cmd     # web应用可执行文件目录，这里存放这服务运行文件和工具等。类似于桌面应用程序的bin目录。
asset   # 存放内置的静态资源目录，不可或缺的静态资源文件。
config  # 站点和应用的配置文件目录，包括json、yaml格式的配置文件。
web     # 前端资源文件目录，页面模板、js、css等，还包含用户的上传文件(upload)。
runtime # 运行时目录，主要有访问、运行、sql等日志，运行时的中间变量等。没必要拿过去，服务启动后时会自动生成对应的目录和文件。
```

<br/>

## 生产使用  
 *（2023年7月17号基于此架构fork的项目首次上线到生产环境，正式使用version版本号，初版号为1.0.0）  
 2023年基于此项目架构给某市政府下辖某委员会开发过一套项目。由于敏感原因这里不贴出具体机构名！  

<br/>

## 最后感谢您参与使用！  
<div>
    <span>银河系 🌌⚛️🧬🧊🔮🗡️✡️🏞️🌈🎮🪞🫧 Requests.</span>
</div>
<div align=right>
    <label>—————— Hyrmaeusj 苏<label>
</div>  

<br/>



