# bandep

Bandep is a tool meant to be used in the CI and/or git commit hooks to prevent
users from adding banned deps.

## Use Cases

- you have multiple packages with the same name in GOPATH (maybe even the same
methods) and your editor sometimes import the wrong import in some files ([example][goreleaser])
- you have explicit rules in your codebase regarding which packages can import
which packages

[goreleaser]: https://github.com/goreleaser/goreleaser/commit/4f3ed001daa737cd4f705e4909bc489936c1a545

## Usage

A basic usage would be to ban some deps in the current project, recursively:

```console
bandep --pkg ./... --ban foo/bar,github.com/foo/bar
```

Or, for example, [on DigitalOcean][do], they forbid the exp package to be
imported from anywhere but the exp package itself:

[do]: https://speakerdeck.com/farslan/go-at-digitalocean?slide=80

```console
go list -f '{{ .Dir }}' -e ./... | grep -v do/exp | while read -r pkg; do
  bandep --pkg $pkg --ban do/exp/foo
done
```

Composing it with shell (and `go list`), you can do basically anything you need.

You can, of course, add it to your git pre-commit hook. Check out
[this example][precom].

[precom]: https://github.com/goreleaser/goreleaser/commit/777b9c68adf5b87a9d3a3291a77a6e41a7215a43

## Install

```console
go get github.com/caarlos0/bandep
```

or

```console
brew install caarlos0/tap/bandep
```

Or download one from the releases tab and install manually.

## How it works

It is quite simple, really: we scan the dirs into an AST, and use that
AST to check if a banned import is being used.
