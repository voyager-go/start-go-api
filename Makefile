# Start Go API - 常用命令

BINARY_NAME=start-go-api
MAIN=main.go

.PHONY: build run test tidy swag docker docker-up docker-down clean help

# 编译
build:
	go build -o $(BINARY_NAME) $(MAIN)

# 运行（需已存在 config.dev.yaml）
run:
	go run $(MAIN) --env dev

# 测试
test:
	go test ./...

# 依赖
tidy:
	go mod tidy

# 生成 Swagger 文档
swag:
	swag init

# Docker 构建镜像
docker:
	docker build -t $(BINARY_NAME):latest .

# Docker Compose 启动
docker-up:
	docker compose up -d

# Docker Compose 停止
docker-down:
	docker compose down

# 清理构建产物
clean:
	rm -f $(BINARY_NAME)

help:
	@echo "Start Go API - 可用目标:"
	@echo "  make build      - 编译二进制"
	@echo "  make run        - 运行（dev 配置）"
	@echo "  make test       - 运行测试"
	@echo "  make tidy       - go mod tidy"
	@echo "  make swag       - 生成 Swagger 文档"
	@echo "  make docker     - 构建 Docker 镜像"
	@echo "  make docker-up  - docker compose 启动"
	@echo "  make docker-down- docker compose 停止"
	@echo "  make clean      - 删除编译产物"
