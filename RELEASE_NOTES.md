# Release: Go 1.26 & Dependency Updates

Tag: release-go-1.26-2026-03-02

Summary
-------
This release upgrades the project to Go 1.26 and refreshes direct and indirect dependencies in conservative batches. It also updates the Docker build to a CGO-enabled multi-stage image to support `github.com/mattn/go-sqlite3`, migrates `mergo` imports to `dario.cat/mergo`, and replaces removed `govc/host/esxcli` usage with an SSH-based executor for ESXi post-configuration.

Notable changes
---------------
- `go.mod`: `go 1.26` and many dependency bumps (see commits).
- Docker: multi-stage CGO-enabled `Dockerfile` and new `docker-compose.yml` for deployment.
- `vmware/esxcli_ssh.go`: New SSH executor replacing removed `govc` esxcli package.
- `api/*`: Updated imports for `dario.cat/mergo`, formatting/import adjustments.

Testing performed
-----------------
- `go build ./...` completed successfully locally.
- Docker image built and smoke-tested locally; `GET /v1/version` returned application version JSON.

Upgrade notes / compatibility
---------------------------
- ESXi provisioning: The app now uses SSH to run equivalent esxcli commands. Ensure ESXi hosts are reachable via SSH and credentials are available in your configuration/secrets.
- If you maintain automated pipelines that expect the old `govc` integration, update them accordingly.

Breaking changes / action required
---------------------------------
- ESXi provisioning now uses SSH execution instead of `govc`/`esxcli`. Ensure target hosts allow SSH access from the service and update stored credentials/secrets accordingly.
- `mergo` moved to a new module path (`dario.cat/mergo`) and was bumped to v1.x; review any custom merge behavior to confirm compatibility.
- Docker build changed to a CGO-enabled multi-stage image; CI pipelines that build the image may need to enable CGO and install build deps or use the provided `Dockerfile`.

How to deploy
-------------
See `docker-compose.yml` for a simple deployment using the local image. For production, push the image to a registry (GHCR, Docker Hub) and update `image:` in the compose file.

Commits & full changelog
------------------------
See the branch: https://github.com/hengelsman/go-via/compare/main...upgrade/go-1.26-deps?expand=1

Credits
-------
Maintainers and contributors who helped test and review dependency upgrades.

Date: 2026-03-02
