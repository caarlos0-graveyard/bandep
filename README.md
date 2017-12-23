# bandep

Bandep is a tool meant to be used in the CI and/or git commit hooks to prevent
users from adding banned deps.

## Use Cases

- you have multiple packages with the same name in GOPATH (maybe even the same
methods) and your editor sometimes import the wrong import in some files
- you have explicit rules in your codebase regarding which packages can import
which packages

## Usage

A basic usage would be to ban some deps in the current project, recursively:

```console
bandep --pkg ./... --ban foo/bar --ban github.com/foo/bar
```

Or, for example, [on DigitalOcean, they forbit the exp package to imported from anywhere][1]:

[1]: https://speakerdeck.com/farslan/go-at-digitalocean?slide=80

```console
go list -f '{{ .Dir }}' -e ./... | grep -v do/exp | while read -r pkg; do
  bandep --pkg $pkg --ban do/exp/foo
done
```

Composing it with shell (and `go list`), you can do basically anything you need.

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
