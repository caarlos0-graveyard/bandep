# bandep

Bandep is a tool meant to be used in the CI and/or git commit hooks to prevent
users from adding banned deps.

## Use Cases

- you have multiple packages with the same name in GOPATH (maybe even the same
methods) and your editor sometimes import the wrong import in some files
- you have explicit rules in your codebase regarding which packages can import
which packages

## Usage

```console
bandep --pkg ./... --ban foo/bar --ban github.com/foo/bar
```

## Install

```console
go get github.com/caarlos0/bandep
```

or

```console
brew install caarlos0/tap/bandep
```

Or download one from the releases tab and install manually.
