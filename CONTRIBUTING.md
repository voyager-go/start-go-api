# 贡献指南

感谢你对 Start Go API 的关注与贡献。

## 如何贡献

- **报告问题**：在 [Issues](https://github.com/voyager-go/start-go-api/issues) 中描述 Bug 或功能建议，并尽量提供复现步骤或环境信息。
- **代码贡献**：通过 Pull Request 提交修改，请先开 Issue 讨论较大改动。

## 开发流程

1. Fork 本仓库并克隆到本地。
2. 创建分支：`git checkout -b feature/xxx` 或 `fix/xxx`。
3. 复制 `config.example.yaml` 为 `config.dev.yaml` 并配置本地环境。
4. 修改代码并确保通过测试：`go test ./...`。
5. 若涉及 API 变更，请更新 Swagger：`swag init`。
6. 提交时写清 commit message，可参考 [Conventional Commits](https://www.conventionalcommits.org/)。
7. 推送到你的 Fork，并针对本仓库的默认分支发起 Pull Request。

## 代码风格

- 遵循 Go 官方 [Effective Go](https://go.dev/doc/effective_go) 与常见格式化（`gofmt` / `goimports`）。
- 新增导出函数/类型请补充注释。
- 保持与现有目录与包划分一致。

## 测试

```bash
go test ./...
```

如有新增逻辑，建议补充单元测试。

## 文档

- 公开 API 请使用 Swagger 注释，并运行 `swag init` 更新文档。
- 配置项、环境变量或运行方式有变更时，请同步更新 README.md。

再次感谢你的参与。
