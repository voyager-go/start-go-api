# Start Go API

基于 [Gin](https://github.com/gin-gonic/gin) 的 Go API 快速开发骨架，适合作为学习 Golang 或二次开发的基础项目。集成 Casbin 权限、JWT 认证、异步任务（Machinery）等常用能力。

[![Go Report Card](https://goreportcard.com/badge/github.com/voyager-go/start-go-api)](https://goreportcard.com/report/github.com/voyager-go/start-go-api)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## 特性

- **Gin**：HTTP 路由与中间件
- **GORM**：MySQL 访问与迁移
- **Redis**：缓存与登录态
- **JWT**：Token 认证
- **Casbin**：RBAC 权限控制
- **Machinery**：基于 Redis 的异步任务
- **Swagger**：API 文档（`swag init` 生成）
- **CLI**：环境/端口等通过命令行参数配置

## 快速开始

### 环境要求

- Go 1.16+
- MySQL 8.0+
- Redis（可选，用于登录态与异步任务）

### 本地运行

```bash
# 克隆项目
git clone https://github.com/voyager-go/start-go-api.git
cd start-go-api

# 安装依赖
go mod tidy

# 复制示例配置并按需修改（数据库、Redis、JWT 密钥等）
cp config.example.yaml config.dev.yaml

# 创建数据库（与 config.dev.yaml 中 mysql.dbname 一致）
# mysql -e "CREATE DATABASE startgoapi CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;"

# 启动（默认端口 8090，使用 config.dev.yaml）
go run main.go
```

访问：

- API：`http://localhost:8090`
- Swagger：`http://localhost:8090/swagger/index.html`（需先执行 `swag init`）

### 命令行参数

```bash
go run main.go -h              # 查看帮助
go run main.go -v              # 查看版本
go run main.go --env dev       # 指定配置：config.dev.yaml
go run main.go --env pre       # 指定配置：config.pre.yaml
go run main.go --port 9999     # 指定端口
```

### 使用 Docker

**仅构建并运行 API 容器**（需已有 MySQL、Redis 或使用下方 compose）：

```bash
cp config.example.yaml config.dev.yaml
# 编辑 config.dev.yaml 中的 mysql / redis 等配置

docker build -t start-go-api .
docker run -p 8090:8090 -v $(pwd)/config.dev.yaml:/app/config.dev.yaml start-go-api --env dev --port 8090
```

**使用 Docker Compose 一键启动 API + MySQL + Redis**：

```bash
cp config.example.yaml config.dev.yaml
# 将 config.dev.yaml 中 mysql.host 改为 mysql，redis.host 改为 redis，并设置 mysql.password: root 等
docker compose up -d
# API: http://localhost:8090
```

首次使用 compose 时，需在应用启动后执行数据库迁移或导入 `docs/sql/startgoapi.sql` 初始化表结构。

## 配置说明

| 配置项 | 说明 |
|--------|------|
| `server.mode` | 运行模式：debug / release |
| `server.jwtSecret` | JWT 签名密钥，生产环境务必使用强随机字符串 |
| `server.tokenExpire` | Token 过期时间（秒） |
| `mysql.*` | MySQL 连接信息 |
| `redis.*` | Redis 连接信息 |

配置文件支持环境变量展开，例如：`password: ${MYSQL_PASSWORD}`。

## 异步任务（Machinery）

项目集成 [Machinery](https://github.com/RichardKnop/machinery)，需在 `main.go` 中开启任务调度，并在 `schedule` 包中编写任务与注册逻辑。

```bash
# 启动 Server：注册任务并投递到 Redis 队列
go run main.go server

# 启动 Worker：消费队列中的任务
go run main.go worker
```

## 生成 API 文档

```bash
swag init
# 或使用 Makefile
make swag
```

## 目录结构

```
├── bootstrap/          # 启动时加载的服务
├── config/             # 配置解析与常量
├── docs/               # Swagger、SQL、请求示例
├── entities/           # 数据模型（含 Casbin）
├── global/             # 全局变量与上下文
├── middleware/         # 中间件（JWT、Casbin、IP 等）
├── modules/            # 业务模块
│   └── system/         # 示例：用户、角色、菜单、API 等
├── pkg/                # 公共包
│   ├── auth/           # JWT
│   ├── lib/            # 日志、MySQL、Redis
│   ├── response/       # 统一响应
│   ├── util/           # 工具函数
│   └── validator/      # 参数校验与翻译
├── repositories/       # 数据访问层
├── router/             # 路由注册
├── schedule/           # 异步任务（Machinery）
├── storage/            # 日志、上传文件等
├── config.example.yaml # 配置示例（复制为 config.dev.yaml 使用）
├── docker-compose.yml
├── Dockerfile
├── main.go
├── Makefile
└── rbac_model.conf     # Casbin 模型
```

## 主要依赖

| 组件 | 说明 |
|------|------|
| [gin](https://github.com/gin-gonic/gin) | HTTP 框架 |
| [gorm](https://github.com/go-gorm/gorm) | ORM |
| [go-redis](https://github.com/go-redis/redis) | Redis 客户端 |
| [casbin](https://github.com/casbin/casbin) | 权限控制 |
| [jwt-go](https://github.com/dgrijalva/jwt-go) | JWT |
| [machinery](https://github.com/RichardKnop/machinery) | 异步任务 |
| [swag](https://github.com/swaggo/swag) | Swagger 文档 |

## 参考项目

- [gf](https://github.com/gogf/gf) / [gf-demos](https://github.com/gogf/gf-demos)
- [go-quick-api](https://github.com/jangozw/go-quick-api)
- [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin)
- [gin-skeleton](https://github.com/mesfreeman/gin-skeleton)

## 贡献

欢迎提交 Issue 与 Pull Request，请先阅读 [CONTRIBUTING.md](CONTRIBUTING.md)。

## 许可证

[MIT](LICENSE)
