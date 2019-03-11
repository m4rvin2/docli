# About

Docli is a Go library for building command-line interfaces. It's developed by [celicoo](https://github.com/celicoo).

## Motivation

I love to write command-line interfaces, but I hate the verbosity envolved - that combined with the lack of maintanance on [docopt](https://github.com/docopt/docopt.go) became my excuse to start Docli.

## How it works

Docli parses the command-line arguments based on the interface described in a doc string:

{{< highlight go >}}
var doc = `
My CLI is cool!

Usage: my-cli [command]

Available Commands:
  init                 initialize the application
  help                 help message

Flags:
  -a, --author string  author name for copyright attribution
  -v, --verbose        verbose output
`
{{< /highlight >}}

## Versioning

Docli uses the [Go Modules](https://github.com/golang/go/wiki/Modules) support built into Go 1.11 and follows the [semver specification](http://semver.org).
