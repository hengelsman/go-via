## Unreleased

- **Go**: Bumped project `go` directive to 1.26.
- **Docker**: Replaced single-file ADD Dockerfile with a CGO-enabled multi-stage build to support `github.com/mattn/go-sqlite3` (see `Dockerfile` and `Dockerfile.local`).
- **Dependencies**: Upgraded numerous direct and indirect dependencies in conservative batches (including `gin`, `gorm` + drivers, `github.com/sirupsen/logrus`, `go-openapi` libs, and many `golang.org/x/*` modules).
- **mergo**: Migrated imports from `github.com/imdario/mergo` to `dario.cat/mergo` and updated to v1.x.
- **VMware/govmomi**: Replaced removed `govc/host/esxcli` usage with an SSH-based executor (`vmware/esxcli_ssh.go`) and adapted `api/postconfig.go` to use SSH execution for ESXi post-configuration.
- **Other**: Applied small formatting/import adjustments across `api/*.go` and updated transitive modules (patch/minor upgrades).

Verification performed:

- Ran `go build ./...` locally — build succeeded.
- Built Docker image and started container; verified startup logs and that `GET /v1/version` returns `{"Version":"dev","Commit":"none","Date":"unknown"}`.

See the PR for full commit history and per-commit notes:
https://github.com/hengelsman/go-via/compare/main...upgrade/go-1.26-deps?expand=1

Date: 2026-03-02
