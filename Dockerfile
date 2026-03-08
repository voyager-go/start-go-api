# 多阶段构建：编译阶段
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 安装 git（go mod 可能依赖）
RUN apk add --no-cache git

# 先复制依赖描述，利用缓存
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 关闭 CGO，静态链接
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -ldflags="-s -w" -o start-go-api .

# 运行阶段
FROM alpine:3.19

WORKDIR /app

# 时区与证书（可选）
RUN apk add --no-cache tzdata ca-certificates && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime 2>/dev/null || true

COPY --from=builder /app/start-go-api .
COPY --from=builder /app/config.example.yaml ./config.example.yaml
COPY --from=builder /app/rbac_model.conf .

# 默认需挂载配置文件或通过环境变量覆盖
# 例如: -v $(pwd)/config.dev.yaml:/app/config.dev.yaml
ENV GIN_MODE=release

EXPOSE 8090

ENTRYPOINT ["./start-go-api"]
CMD ["--env", "dev", "--port", "8090"]
