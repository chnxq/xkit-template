# xkit-template

`xkit-template` is the starter service template for `xkit`. It is not a demo project. It is a runnable project skeleton that can be copied into a target repository and then maintained as normal application code.

The template owns stable project structure and startup flow. `xkit` owns generated resource code derived from API, proto, Ent schema, and generation config.

Chinese version: [README.md](README.md)

## What It Contains

- `cmd/server`: service command entry.
- `cmd/server/assets`: startup assets such as OpenAPI and RBAC files.
- `configs`: default configuration files.
- `internal/bootstrap`: config loading, factory preloading, logging, registry, tracing, and app assembly.
- `internal/server`: HTTP/REST, gRPC, Asynq, and SSE transport construction plus registration hooks.
- `internal/data/bootstrap`: extension point for shared data providers.
- `template.yaml`: template variables, ignore rules, preserved files, and obsolete files.

## Recommended Usage

Run from the `xkit` repository:

```powershell
go run ./cmd/xkit init template D:\GoProjects\XAdmin\xkit-template `
  --project D:\GoProjects\XAdmin\xadmin-web `
  --module xadmin-web `
  --app-name XAdmin `
  --command-name xadmin-web `
  --service-name admin
```

After copying files, `xkit init template` runs this in the target project:

```powershell
go get -u all
```

To preview file writes:

```powershell
go run ./cmd/xkit init template D:\GoProjects\XAdmin\xkit-template `
  --project D:\GoProjects\XAdmin\xadmin-web `
  --module xadmin-web `
  --app-name XAdmin `
  --command-name xadmin-web `
  --service-name admin `
  --dry-run
```

To skip dependency updates:

```powershell
go run ./cmd/xkit init template D:\GoProjects\XAdmin\xkit-template `
  --project D:\GoProjects\XAdmin\xadmin-web `
  --module xadmin-web `
  --app-name XAdmin `
  --command-name xadmin-web `
  --service-name admin `
  --skip-go-get-update-all
```

Then generate resource code:

```powershell
go run ./cmd/xkit gen all admin `
  --project D:\GoProjects\XAdmin\xadmin-web `
  --config D:\GoProjects\XAdmin\xkit\examples\xadmin-web\admin.yaml
```

## Ownership Boundary

The template owns the stable startup skeleton:

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
- `internal/server/http_options.go`
- `internal/server/grpc_options.go`
- `internal/server/asynq_options.go`
- `internal/server/sse_options.go`
- `internal/server/transport_config.go`
- `internal/server/options.go`
- `internal/server/sse.go`
- `internal/server/tls.go`
- `internal/data/bootstrap/data.go`
- `internal/data/bootstrap/resources.go`
- `configs/*.yaml`

`xkit gen all` should only overwrite resource-related `*.gen.go` files, for example:

- `internal/bootstrap/generated_servers.gen.go`
- `internal/data/bootstrap/ent_client.gen.go`
- `internal/server/rest_register.gen.go`
- `internal/server/grpc_register.gen.go`
- `internal/service/*_service.gen.go`
- `internal/data/repo/*_repo.gen.go`

## Manual Extension Points

These files are intended for project-specific code and are preserved during template sync:

- `internal/bootstrap/hooks.go`: extra transport servers, lifecycle hooks, or background jobs.
- `internal/server/options.go`: project-specific HTTP/gRPC business middleware.
- `internal/data/bootstrap/data.go`: shared data providers such as Redis, cache, OSS, or queues.
- `internal/data/bootstrap/resources.go`: lifecycle entry for shared data resources.
- `internal/service/*_service_ext.go`: handwritten service extensions.
- `internal/data/repo/*_repo_ext.go`: handwritten repository extensions.

## Template Rules

In `template.yaml`:

- `variables` defines default values to replace during initialization.
- `ignore` defines files that are not copied.
- `preserve` defines existing target files that should be kept.
- `obsolete` defines old template files removed during `--force` sync.

Currently obsolete files:

- `cmd/server/wire.go`
- `cmd/server/wire_gen.go`

## Local Verification

The template itself should stay buildable:

```powershell
go test ./...
```

After applying the template and generating resources in a target project, verify with:

```powershell
go test ./...
go run ./cmd/server server -config_path ./configs
```
