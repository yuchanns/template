# Go 项目结构实践
本仓库采用了目前社区中较为流行的 Gin, Gorm, Dig, Go-Resty 等第三方库作为使用案例，但理论上可以结合各种具有良好设计规范的工具包使用。
## 项目组织结构概览
文件树概览，后续会对照解说，此时不需要纠结。
```sh
.
├── go.mod
├── go.sum
├── internal
│   ├── domain
│   │   ├── async
│   │   │   ├── domain.go
│   │   │   └── repo.go
│   │   └── greet
│   │       ├── domain.go
│   │       └── repo.go
│   ├── infra
│   │   ├── client
│   │   │   └── unknown.go
│   │   ├── mysql
│   │   │   ├── greet.go
│   │   │   └── models
│   │   │       └── greeter.go
│   │   └── redis
│   └── server
│       ├── async.go
│       └── greeting.go
├── main.go
├── startup
│   ├── di.go
│   ├── event.go
│   ├── route.go
│   └── vars.go
├── utils
│   ├── converter.go
│   ├── gin_wrapper.go
│   ├── goroutine.go
│   └── ioc.go
└── vars
    └── var.go
```
![](./structure.png)
### 分层与单向依赖
Go 的一个特色同时也是缺陷：不支持循环引用。
> 1. 为了高速编译(主要原因)
> 2. 单向依赖逻辑简洁(大道至简)

所以我们为了避免在应用变得复杂后踩到循环引用的坑，一开始代码结构就要约定好**单边方向**，如上图。

Server 层也就是 PHP 中 Laravel 的 Controller 层，只负责获取入参、校验和调用 Domain 方法进行数据处理。

Domain 层用于编写具体而复杂的业务逻辑。
> 有的团队会把 Domain 细分成 Service 和 Domain 或者 App 和 Service 或者三者皆有。实际开发过程中发觉这样太过复杂，不需要分那么多层，所以一般取一个 Domain 层就足矣。

同时可以注意到 Domain 还包含了一些以 Repo 为名字后缀的接口，这是用来定义数据交互的接口，对业务逻辑屏蔽了具体的细节和差异。 Domain 业务逻辑通过调用这些接口来与实际的数据库、缓存或者远程调用进行交互。
> 使用接口编程的好处之一是和具体实现解绑。
>
> 随着业务发展我们可能会出现将本地数据库调用拆分出去成为远程调用、觉得关系型数据库不符合业务场景需要切换成非关系型数据库等场景，这时候我们变更实现会更清晰和容易。
>
> 严谨地说，对于 Server 层和 Domain 层也应该定义接口，面向接口编程。但是这样写代码过于繁琐，最好省略。

Infra 层是基础设施层，真正进行数据交互操作的地方，实现了在 Domain 层定义的 Repo 接口。封装了所有 sql 细节、缓存交互细节和远程调用细节。
> 通常开发者在操作 mysql 数据库时可以很好的隔绝开业务代码和 sql 做到不互相侵入。但是遇到类似 mongodb 这类没有足够抽象的 orm 支持的数据库，就会出于代码复用等原因不自觉地开始把 sql 细节混入业务逻辑中，造成混乱。
> 
> 用 Infra 层可以使这种混乱具有边界。

整个调用方向是 `Server -> Domain -> Infra` ，不存在反向调用。

> 有时候我们确实会存在每个层都需要使用同一个实例的场景，例如：
> 1. 数据结构转换，入参 request 数据结构转换成 domain 层的 DTO(Data Transfer Object) 、DTO 转换成具体数据库的 DO(Data Object)。
> 2. Channel 通道，在 Domain 层投递数据，在 Server 层取数据。
> 3. 数据库实例，在 Server 层创建，在 Infra 层使用。
> 
> 这种时候建议把代码写在不属于任何一层的**第三方目录**，例如根目录的 common 、utils 等目录下。这样任一层都可以调用而不会出现循环引用编译失败。

### 目录划分与代码细节
