# xkit-template

`xkit-template` 是 `xkit` 的服务项目起始模板。它不是演示项目，而是一个可以被复制到目标项目后继续演进的启动骨架。

模板负责稳定的项目结构和启动链路，`xkit` 负责根据 API、proto、Ent schema 和生成配置写入资源相关代码。

English version: [README.en.md](README.en.md)

## 包含内容

- `cmd/server`：服务命令入口。
- `cmd/server/assets`：OpenAPI、RBAC 等启动资产。
- `configs`：默认配置文件。
- `internal/bootstrap`：配置加载、工厂预加载、日志、注册中心、trace 和应用装配入口。
- `internal/server`：HTTP/gRPC server 构造和注册挂点。
- `internal/data/bootstrap`：数据层公共 provider 预留位置。
- `template.yaml`：模板变量、忽略规则、保留文件和废弃文件清单。

## 推荐使用方式

在 `xkit` 仓库中执行：

```powershell
go run ./cmd/xkit init template D:\GoProjects\XAdmin\xkit-template `
  --project D:\GoProjects\XAdmin\xadmin-web `
  --module xadmin-web `
  --app-name XAdmin `
  --command-name xadmin-web `
  --service-name admin
```

真实复制模板后，`xkit init template` 会自动在目标项目执行：

```powershell
go get -u all
```

如果只想预览写入计划：

```powershell
go run ./cmd/xkit init template D:\GoProjects\XAdmin\xkit-template `
  --project D:\GoProjects\XAdmin\xadmin-web `
  --module xadmin-web `
  --app-name XAdmin `
  --command-name xadmin-web `
  --service-name admin `
  --dry-run
```

如果处于离线环境或不希望更新依赖：

```powershell
go run ./cmd/xkit init template D:\GoProjects\XAdmin\xkit-template `
  --project D:\GoProjects\XAdmin\xadmin-web `
  --module xadmin-web `
  --app-name XAdmin `
  --command-name xadmin-web `
  --service-name admin `
  --skip-go-get-update-all
```

初始化后再运行资源代码生成：

```powershell
go run ./cmd/xkit gen all admin `
  --project D:\GoProjects\XAdmin\xadmin-web `
  --config D:\GoProjects\XAdmin\xkit\examples\xadmin-web\admin.yaml
```

## 生成边界

模板拥有固定启动骨架，后续可由项目手写维护：

- `cmd/server/main.go`
- `cmd/server/server.go`
- `internal/bootstrap/app.go`
- `internal/bootstrap/cleanup.go`
- `internal/bootstrap/factories.go`
- `internal/bootstrap/infra.go`
- `internal/bootstrap/hooks.go`
- `internal/server/asynq.go`
- `internal/server/http.go`
- `internal/server/grpc.go`
- `internal/server/options.go`
- `internal/server/sse.go`
- `internal/server/tls.go`
- `internal/data/bootstrap/data.go`
- `internal/data/bootstrap/resources.go`
- `configs/*.yaml`

`xkit gen all` 只应该覆盖资源相关的 `*.gen.go` 文件，例如：

- `internal/bootstrap/generated_servers.gen.go`
- `internal/data/bootstrap/ent_client.gen.go`
- `internal/server/rest_register.gen.go`
- `internal/server/grpc_register.gen.go`
- `internal/service/*_service.gen.go`
- `internal/data/repo/*_repo.gen.go`

## 手写扩展点

这些文件用于后续项目定制，并且模板同步时默认保留：

- `internal/bootstrap/hooks.go`：额外 transport server、生命周期或后台任务。
- `internal/server/options.go`：项目业务侧 HTTP/gRPC middleware 挂点。
- `internal/data/bootstrap/data.go`：Redis、缓存、OSS、队列等公共数据 provider。
- `internal/data/bootstrap/resources.go`：公共数据资源生命周期入口。
- `internal/service/*_service_ext.go`：业务服务手写扩展。
- `internal/data/repo/*_repo_ext.go`：数据访问手写扩展。

## 模板规则

`template.yaml` 中：

- `variables` 定义初始化时会替换的默认值。
- `ignore` 定义不会复制的文件。
- `preserve` 定义目标项目已有时应保留的文件。
- `obsolete` 定义 `--force` 同步时需要清理的旧模板文件。

当前会清理的废弃文件：

- `cmd/server/wire.go`
- `cmd/server/wire_gen.go`

## 本地验证

模板自身应保持可编译：

```powershell
go test ./...
```

目标项目完成模板初始化和资源生成后，也应执行：

```powershell
go test ./...
go run ./cmd/server server -config_path ./configs
```
