# goflake

Start a Go CLI project with a pinned development environment, a small Cobra
entry point, and consistent build and lint commands.

- **One template:** `nix flake init -t github:kevinpita/goflake`.
- **Complete development shell:** Go, editor/debugger tools, Just, and native CGO tools.
- **Separate layers:** `cmd/app -> internal/cli -> internal/app`.
- **Automatic fixes:** gofumpt, goimports, and supported linter fixes, with a JSON report.
- **No example tests:** add tests for your project's actual behavior.

## Create a project

```sh
mkdir -p ~/code/myproject
cd ~/code/myproject
nix flake init -t github:kevinpita/goflake
devenv shell
just init github.com/you/myproject
git init
just check
just run --name Go
```

The example prints `Hello, Go!`. No clone of goflake is needed.

Run `just init` once, before development. It validates your module path and
updates `go.mod`, both internal imports, and the goimports local prefix.

> [!NOTE]
> The flake supplies the starter files. Generated projects use `devenv shell`,
> not `nix develop`. Optional direnv support is included. Run `direnv allow`
> after inspecting `.envrc` if you want automatic shell loading.

## Project layout

```text
cmd/app/       -> executable entry point and exit status
internal/cli/  -> Cobra commands, flags, and output
internal/app/  -> application behavior, independent of Cobra
scripts/       -> module setup
```

The CLI returns errors to the entry point and uses caller-controlled output
writers. There is no plugin framework or Viper dependency. Replace the greeting
example with your application logic and add tests for its boundaries and failures.

## Daily commands

| Command | Action |
| --- | --- |
| `just check` | Lint/fix, test, race-test, and build |
| `just lint` | Apply fixes and write terminal text plus a JSON report |
| `just run --name Go` | Run the starter CLI |
| `just build` | Build `bin/app` |
| `just test` / `just race` | Run your future tests |
| `just` | List all recipes |

See the [generated project README](template/README.md) for coverage, formatting,
module maintenance, and cleanup commands.

> [!IMPORTANT]
> Lint can change source files, `go.mod`, and `go.sum`. Review the changes before
> committing. Remaining findings cause a nonzero exit code. The latest report,
> `.golangci-lint-report.json`, is ignored by Git.

## Tools

| Area | Included |
| --- | --- |
| Go | Go 1.27, gopls, Delve, staticcheck, editor/generator tools |
| Checks | golangci-lint, gofumpt and goimports through golangci-lint |
| Native builds | C/C++ compiler, pkg-config, Make, CMake, GDB on Linux |
| Commands | Just, Bash, core shell utilities, Git, curl, jq, nixfmt |

`devenv.lock` pins package revisions. CGO is enabled for native builds and race
tests. Add project-specific C libraries to `packages` in `devenv.nix`.
Template generation, lint, builds, and CGO compilation have been verified on
`x86_64-linux`.

## Maintain the defaults

Edit files in `template/` to change future projects. Existing projects are
independent copies and do not change automatically. Track new template files
with `git add template` so Nix includes them when creating projects.

From an empty project directory, test a local checkout with:

```sh
nix flake init -t "$HOME/code/goflake"
```

Do not add a `path:` prefix. The Git source excludes ignored build output and
devenv state.

To update the pinned tools:

```sh
cd ~/code/goflake/template
devenv update
devenv shell -- just check
```

Review and commit `template/devenv.lock` with the configuration. Keep the Go
version in `template/devenv.nix` and `template/go.mod` aligned when upgrading Go.
