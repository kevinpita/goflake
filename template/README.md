# Go project

A Go starter with devenv, Just, golangci-lint, and CGO support.
The starter module path is `example.com/project`.

## Start

From the project directory:

```sh
devenv shell
just init github.com/you/project
just check
```

Run `just init` once, before development. It updates `go.mod`, the internal
import paths, and the goimports local prefix. It validates the module path first.
Initialize Git with `git init` if this directory is not already a repository.

For automatic shell loading, run `direnv allow`.
Otherwise, use `devenv shell` or `devenv shell -- just check`.
This project uses native devenv files, not `nix develop`.

## Commands

| Command | Action |
| --- | --- |
| `just` | List recipes |
| `just check` | Lint/fix, test, race-test, and build |
| `just lint` | Apply supported fixes and report remaining findings |
| `just format` | Run gofumpt and goimports through golangci-lint |
| `just format-nix` | Format `devenv.nix` |
| `just test` | Run Go tests |
| `just race` | Run the race detector |
| `just coverage` | Write and summarize `coverage.out` |
| `just build` | Build `bin/app` |
| `just run --name Go` | Run the Cobra CLI with a name |
| `just tidy` | Update module files for the current imports |
| `just clean` | Remove build output and reports |

> [!IMPORTANT]
> Lint applies supported fixes and can modify source files, `go.mod`, and `go.sum`.
> Review changes before committing. Remaining findings cause a nonzero exit code.

Lint produces terminal text and `.golangci-lint-report.json` in one run.
The report is ignored by Git and replaced on the next run.

## Layout

```text
cmd/app/       -> executable entry point and exit status
internal/cli/  -> Cobra commands, flags, and output
internal/app/  -> application behavior
scripts/       -> project setup
```

Dependencies flow `cmd/app -> internal/cli -> internal/app`. The core does not
import Cobra. The CLI uses command writers to keep output under caller control
and returns errors to the entry point without printing duplicate errors or usage.

No example tests are included. Add tests for your project's behavior, boundaries,
and failure cases. The test, race, and coverage recipes are ready when needed.

The starter uses Cobra and the Go standard library. Run `just run --help` for
help. Add commands and rename `app` when the project needs them.

## Development tools

The locked environment supplies Go 1.27, Just, golangci-lint, gopls, Delve,
staticcheck, goimports, and Go editor/generator tools. It also supplies a native
C/C++ compiler, pkg-config, Make, CMake, Bash, core shell utilities, Git, curl,
jq, and nixfmt. Linux environments include GDB.

CGO is enabled. Nix supplies the native compiler wrappers and their
header/library paths through the shell.
Add project-specific C libraries to `packages` in `devenv.nix`. For example,
a project that links SQLite must also add `pkgs.sqlite`.
To make a separate pure-Go build, use `CGO_ENABLED=0 go build ./cmd/app`.
The race detector still requires CGO.

## Update tools

Commit `devenv.lock` so contributors use the same package revisions.
Run `devenv update`, then `just check`, to update and verify packages.
Go's minor version is explicit in `devenv.nix` and `go.mod`. Change both when
upgrading Go. Use `devenv.local.nix` only for local settings.
