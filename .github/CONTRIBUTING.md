# Contributing to Objx

Thanks for your interest in improving objx! Issues and pull requests are welcome.

## Before you start

- For bug fixes and small improvements, feel free to open a pull request directly.
- For new features or API changes, please open an issue first so we can discuss
  the approach before you spend time on it.
- Security vulnerabilities should **not** be reported in public issues. See
  [SECURITY.md](../SECURITY.md) instead.

## Development setup

You need a supported Go version (the three most recent major releases) and,
optionally, [Task](https://taskfile.dev) and
[golangci-lint](https://golangci-lint.run).

```sh
task test      # go test -race ./...
task lint      # golangci-lint run ./...
task generate  # go generate ./...
```

Each task is a thin wrapper, so you can run the underlying commands directly.

## Generated code

The type-specific accessors in `type_specific_codegen.go` and their tests in
`type_specific_codegen_test.go` are generated from the templates in `codegen/`.
Do not edit the generated files by hand. Change the templates and run
`task generate` instead. CI fails if the generated files are out of date.

## Pull requests

- Add or update tests for any change in behavior.
- Make sure `task test` and `task lint` pass locally. CI runs the same checks
  against all supported Go versions.
- Avoid language features or standard library APIs newer than the minimum Go
  version in `go.mod`.
- Keep pull requests focused. Unrelated changes are easier to review separately.

By contributing, you agree to abide by our [Code of Conduct](CODE_OF_CONDUCT.md).
