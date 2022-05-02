# Chalk - Terminal string styling

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/chalk)](https://pkg.go.dev/github.com/go-zoox/chalk)
[![Build Status](https://github.com/go-zoox/chalk/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/chalk/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/chalk)](https://goreportcard.com/report/github.com/go-zoox/chalk)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/chalk/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/chalk?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/chalk.svg)](https://github.com/go-zoox/chalk/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/chalk.svg?label=Release)](https://github.com/go-zoox/chalk/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/chalk
```

## Getting Started

```go
import (
  "testing"
  "github.com/go-zoox/chalk"
)

func main() {
	fmt.Println(chalk.Red("Hello, World!"))
}
```

## Inspired by
* [chalk/chalk](https://github.com/chalk/chalk) - 🖍 Terminal string styling done right
* [ttacon/chalk](https://github.com/ttacon/chalk) - Intuitive package for prettifying terminal/console output

## License
GoZoox is released under the [MIT License](./LICENSE).
