### 项目名称 
#### Start Go Api

### 项目用途
#### 使用gin框架作为基础开发库，封装一套适用于面向api编程的快速开发结构，适合刚学习Golang的新手

### 生成API文档
```shell
swag init # 默认使用swagger作为文档管理工具 
```
### 项目启动方式
```shell
go mod tidy
SERVER_ENV=dev go run main.go # dev是默认开发环境，请自行修改根目录下dev和pre对应的配置文件
```

### 目录结构介绍
```shell
|-api             # 接口存放目录
|-bootstap        # 程序启动时需要加载的服务
|-config          # 解析配置文件
|-dao             # 数据库的增删改查
|-docs            # 存放一些swagger接口文档与api请求示例以及SQL文件
  |-request_demo  # jetbrains自带的HTTP请求示例
  |-sql           # 项目初始化时的SQL参考示例
|-entity          # 存放表对应的实体，可以理解为model
|-global          # 一些全局变量以及全局方法
|-middleware      # 实现简单的中间件
|-pkg             # 自定义的常用服务、JWT、助手函数与格式化返回
  |-auth          # jwt认证
  |-lib           # 构造日志服务、数据库服务、Redis服务等
  |-response      # 返回值的格式化处理
  |-util          # 助手函数
|-router          # 路由注册
|-service         # 接口的具体业务处理
|-storage         # 默认存放一些资源文件，如日志文件、上传文件等
```
### 基础组件

- gin框架     项目地址: https://github.com/gin-gonic/gin
- yaml配置库 项目地址: https://github.com/go-yaml/yaml
- logrus日志库以及相关钩子
    - https://github.com/sirupsen/logrus
    - https://github.com/rifflock/lfshook
    - https://github.com/lestrrat-go/file-rotatelogs
- gorm v2 对象关联映射库 项目地址: https://github.com/go-gorm/gorm
- strutil 常用工具库 项目地址: https://github.com/gookit/goutil/tree/master/strutil
- go-redis 项目地址: https://github.com/go-redis/redis
- jwt-go 项目地址: https://github.com/dgrijalva/jwt-go
- cors 项目地址: https://github.com/gin-contrib/cors
- gconv 类型转换库 项目地址: https://godoc.org/github.com/gogf/gf/util/gconv

### 项目在编写与设计时参考了Github上一些优秀的项目
- https://github.com/gogf/gf
- https://github.com/gogf/gf-demos
- https://github.com/jangozw/go-quick-api
- https://github.com/flipped-aurora/gin-vue-admin
- https://github.com/mesfreeman/gin-skeleton